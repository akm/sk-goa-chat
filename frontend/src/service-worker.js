/// <reference types="@sveltejs/kit" />
// See https://kit.svelte.jp/docs/service-workers
// import { build, files, version } from '$service-worker';

// すべてのリクエストに ID トークンを含める
// See https://firebase.google.com/docs/auth/web/service-worker-sessions?hl=ja

// import { initializeApp } from 'firebase/app';
// import { getAuth, } from 'firebase/auth';
import { getIdToken, onAuthStateChanged } from 'firebase/auth';
// import { firebaseConfig } from '$lib/firebase/firebaseconfig';

import { auth } from '$lib/firebase/auth';

auth.onAuthStateChanged((user) => {
	if (!user) {
		console.log('service-worker/auth.onAuthStateChanged no user');
		return;
	}
	console.log('service-worker/auth.onAuthStateChanged user.uid', user.uid);

	// https://developer.mozilla.org/ja/docs/Web/API/Clients/claim
	self.clients.claim().then(() => {
		// https://developer.mozilla.org/ja/docs/Web/API/Clients
		self.clients.matchAll().then((clients) => {
			console.log('service-worker/auth.onAuthStateChanged clients', clients);

			// https://developer.mozilla.org/ja/docs/Web/API/Client/postMessage
			clients.forEach((client) => client.postMessage({ uid: user.uid }));
		});
	});
});

// Types:
//
// self は ServiceWorkerGlobalScope のインスタンスっぽい
// ServiceWorkerGlobalScope --|> WorkerGlobalScope --|> EventTarget
// - https://developer.mozilla.org/ja/docs/Web/API/ServiceWorkerGlobalScope
// - https://developer.mozilla.org/ja/docs/Web/API/WorkerGlobalScope
// - https://developer.mozilla.org/ja/docs/Web/API/EventTarget
//
// FetchEvent: https://developer.mozilla.org/ja/docs/Web/API/FetchEvent
self.addEventListener('fetch', (event) => {
	console.log('service-worker/fetch event:', event);

	/** @type {FetchEvent} */
	const evt = event;

	const requestProcessor = (idToken) => {
		console.log('service-worker/fetch/requestProcessor idToken:', idToken);

		let req = evt.request;
		let processRequestPromise; //: Promise<void>
		// For same origin https requests, append idToken to header.
		if (
			self.location.origin == getOriginFromUrl(evt.request.url) &&
			(self.location.protocol == 'https:' || self.location.hostname == 'localhost') &&
			idToken
		) {
			// Clone headers as request headers are immutable.
			const headers = new Headers();
			req.headers.forEach((val, key) => {
				headers.append(key, val);
			});
			// Add ID token to header.
			// headers.append('Authorization', 'Bearer ' + idToken);
			headers.append('X-ID-TOKEN', idToken || '');
			processRequestPromise = getBodyContent(req).then((body) => {
				try {
					req = new Request(req.url, {
						method: req.method,
						headers: headers,
						mode: 'same-origin',
						credentials: req.credentials,
						cache: req.cache,
						redirect: req.redirect,
						referrer: req.referrer,
						body
						// bodyUsed: req.bodyUsed,
						// context: req.context
					});
				} catch (e) {
					// This will fail for CORS requests. We just continue with the
					// fetch caching logic below and do not pass the ID token.
				}
			});
		} else {
			processRequestPromise = Promise.resolve();
		}
		return processRequestPromise.then(() => {
			return fetch(req);
		});
	};
	// Fetch the resource after checking for the ID token.
	// This can also be integrated with existing logic to serve cached files
	// in offline mode.
	evt.respondWith(getIdTokenPromise().then(requestProcessor, requestProcessor));
});

const getOriginFromUrl = (url) => {
	// https://stackoverflow.com/questions/1420881/how-to-extract-base-url-from-a-string-in-javascript
	const pathArray = url.split('/');
	const protocol = pathArray[0];
	const host = pathArray[2];
	return protocol + '//' + host;
};

// Get underlying body if available. Works for text and json bodies.
const getBodyContent = (req) => {
	return Promise.resolve()
		.then(() => {
			if (req.method !== 'GET') {
				if (req.headers.get('Content-Type').indexOf('json') !== -1) {
					return req.json().then((json) => {
						return JSON.stringify(json);
					});
				} else {
					return req.text();
				}
			}
		})
		.catch((error) => {
			console.error('Request body processing failed: ', error);
		});
};

/**
 * Returns a promise that resolves with an ID token if available.
 * @return {!Promise<?string>} The promise that resolves with an ID token if
 *     available. Otherwise, the promise resolves with null.
 */
const getIdTokenPromise = () => {
	return new Promise((resolve, reject) => {
		const unsubscribe = onAuthStateChanged(auth, (user) => {
			console.log('service-worker/getIdTokenPromise/onAuthStateChanged user:', user);
			unsubscribe();
			if (user) {
				getIdToken(user).then(
					(idToken) => {
						console.log(
							'service-worker/getIdTokenPromise/onAuthStateChanged/getIdToken idToken:',
							idToken
						);
						resolve(idToken);
					},
					(error) => {
						reject(error);
					}
				);
			} else {
				resolve(null);
			}
		});
	});
};
