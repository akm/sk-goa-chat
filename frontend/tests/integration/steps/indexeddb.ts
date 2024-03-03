import fs from 'fs';
import { test } from '@playwright/test';
import type { Page } from '@playwright/test';

// IndexedDB については、以下のページを参考にしました。
// https://developer.mozilla.org/ja/docs/Web/API/IndexedDB_API/Using_IndexedDB

export const saveIndexedDBDataTo = async (
	page: Page,
	dbName: string,
	objectStoreName: string,
	filePath: string
): Promise<void> => {
	await test.step(`save indexed DB ${dbName}.${objectStoreName} to ${filePath}`, async () => {
		const data = await page.evaluate(exportIndexedDBStore, { dbName, objectStoreName });
		fs.writeFileSync(filePath, JSON.stringify(data), 'utf-8');
	});
};

// この関数は Page.evaluate を使ってブラウザのコンテキストで実行されるので、関数を分割しないようにしてください。
const exportIndexedDBStore = (arg: {
	dbName: string;
	objectStoreName: string;
}): Promise<string> => {
	const indexedDB = window.indexedDB;
	const { dbName, objectStoreName } = arg;
	return new Promise((resolve, reject) => {
		console.log('indexedDB.open is calling');
		const request = indexedDB.open(dbName);
		request.onerror = (event: Event) => {
			console.log('indexedDB.open request.onerror', event);
			reject(new Error(`failed to open ${dbName} ${event.target}`));
		};
		request.onsuccess = (event: Event) => {
			console.log('indexedDB.open request.oncussess', event);
			const db: IDBDatabase = (event.target as IDBRequest).result;
			db.onerror = (event: Event) => {
				console.log('db.onerror', event);
				reject(new Error(`Database error ${dbName} ${event.target}`));
			};
			// let result: string | undefined;
			// db.onclose = (event: Event) => {
			// 	console.log(`db.onclose success: ${result}`, result);
			// 	if (result) {
			// 		resolve(result);
			// 	} else {
			// 		reject(new Error(`Database closed ${dbName} ${event.target}`));
			// 	}
			// };

			const tx = db.transaction([objectStoreName]);
			const store = tx.objectStore(objectStoreName);

			const req = store.getAll();
			req.onerror = (event: Event) => {
				console.log('req.onerror', event);
				reject(new Error(`failed to get all in ${objectStoreName} ${event.target}`));
			};
			req.onsuccess = (event: Event) => {
				console.log(`objectStore.getAll req.onsuccess`, event);
				const result = JSON.stringify(req.result);
				db.onclose = (event: Event) => {
					console.log(`db.onclose`, event);
					resolve(result);
				};
				console.log(`objectStore.getAll req.onsuccess db closing`);
				db.close();
				console.log(`objectStore.getAll req.onsuccess db closing done`);
				// resolve(result);
			};
		};
	});
};

export const loadIndexedDBDataFrom = async (
	page: Page,
	dbName: string,
	objectStoreName: string,
	keyPath: string,
	filePath: string
): Promise<void> => {
	await test.step(`load indexed DB ${dbName}.${objectStoreName} from ${filePath}`, async () => {
		const data = fs.readFileSync(filePath, 'utf-8');
		await page.evaluate(importIndexedDBStore, { dbName, objectStoreName, keyPath, data });
	});
};

// この関数は Page.evaluate を使ってブラウザのコンテキストで実行されるので、関数を分割しないようにしてください。
const importIndexedDBStore = (arg: {
	dbName: string;
	objectStoreName: string;
	keyPath: string;
	data: string;
}): Promise<void> => {
	const indexedDB = window.indexedDB;
	const { dbName, objectStoreName, keyPath, data } = arg;
	return new Promise((resolve, reject) => {
		let objects: any; // eslint-disable-line @typescript-eslint/no-explicit-any
		try {
			objects = JSON.parse(data);
		} catch (e) {
			reject(new Error(`failed to parse JSON ${arg.data}`));
		}

		console.log('indexedDB.open is calling');
		const request = indexedDB.open(dbName, 1);
		request.onerror = (event: Event) => {
			console.log('indexedDB.open request.onerror', event);
			reject(new Error(`failed to open ${dbName} ${event.target}`));
		};
		request.onupgradeneeded = (event) => {
			console.log('indexedDB.open request.onupgradeneeded', event);
			const db: IDBDatabase = (event.target as IDBRequest).result;
			// let success = false;
			// db.onclose = (event: Event) => {
			// 	console.log(`db.onclose success: ${success}`, event);
			// 	if (success) {
			// 		resolve();
			// 	} else {
			// 		reject(new Error(`Database closed ${dbName} ${event.target}`));
			// 	}
			// };

			const objectStore = db.createObjectStore(objectStoreName, { keyPath });

			// データを追加する前に objectStore の作成を完了させるため、
			// transaction oncomplete を使用します。
			objectStore.transaction.oncomplete = (event: Event) => {
				console.log('objectStore.transaction.oncomplete', event);
				// 新たに作成した objectStore に値を保存します。
				const customerObjectStore = db
					.transaction(objectStoreName, 'readwrite')
					.objectStore(objectStoreName);
				for (const obj of objects) {
					customerObjectStore.add(obj);
				}
				// success = true;
				console.log('objectStore.transaction.oncomplete db closing');
				db.close();
				console.log('objectStore.transaction.oncomplete db closing done');
			};
		};
	});
};
