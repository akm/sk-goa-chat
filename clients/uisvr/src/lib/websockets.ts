import 'websocket-polyfill';

let _notificationsSocket: WebSocket;

// 呼び出す方が authReady を確認すること
export const notificationsSocket = (idToken: string): Promise<WebSocket> => {
	console.log('notificationsSocket with idToken', idToken);

	const rproxyOrigin = import.meta.env.VITE_RPROXY_ORIGIN as string;
	const wsOrigin = rproxyOrigin.includes('https://')
		? rproxyOrigin.replace('https://', 'wss://')
		: rproxyOrigin.replace('http://', 'ws://');
	const wsUrl = `${wsOrigin}/ws/notifications/subscribe?token=${idToken}`;

	console.log('notificationsSocket wsUrl', wsUrl);

	return new Promise((resolve, reject) => {
		// const websocket = io(wsUrl);
		const ws = new WebSocket(wsUrl);
		ws.addEventListener('open', (event) => {
			console.log('WebSocket open', event);
			resolve(ws);
		});
		ws.addEventListener('close', (event) => {
			console.log('WebSocket close', event);
		});
		ws.addEventListener('error', (event) => {
			console.log('WebSocket error', event);
			reject(new Error(`WebSocket error: ${event}`));
		});

		console.log('notificationsSocket ws', ws);

		_notificationsSocket = ws;
		// return _notificationsSocket;
	});
};

export const closeWebSockets = (): void => {
	console.log('closeWebSockets _notificationsSocket', _notificationsSocket);
	if (!_notificationsSocket) return;
	console.log('closeWebSockets closing');
	_notificationsSocket.close(1000);
	console.log('closeWebSockets closing done');
};
