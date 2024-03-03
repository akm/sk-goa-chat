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
	const { dbName, objectStoreName } = arg;
	return new Promise((resolve, reject) => {
		const request = indexedDB.open(dbName);
		request.onerror = (event: Event) => {
			reject(new Error(`failed to open ${dbName} ${event.target}`));
		};
		request.onsuccess = (event: Event) => {
			const db: IDBDatabase = (event.target as IDBRequest).result;
			db.onerror = (event: Event) => {
				reject(new Error(`Database error ${dbName} ${event.target}`));
			};

			const tx = db.transaction([objectStoreName]);
			const store = tx.objectStore(objectStoreName);

			const req = store.getAll();
			req.onerror = (event: Event) => {
				reject(new Error(`failed to get all in ${objectStoreName} ${event.target}`));
			};
			req.onsuccess = () => {
				resolve(JSON.stringify(req.result));
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
		const data = JSON.parse(fs.readFileSync(filePath, 'utf-8'));
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
	const { dbName, objectStoreName, keyPath } = arg;
	return new Promise((resolve, reject) => {
		let obj: any; // eslint-disable-line @typescript-eslint/no-explicit-any
		try {
			obj = JSON.parse(arg.data);
		} catch (e) {
			reject(new Error(`failed to parse JSON ${arg.data}`));
		}

		const request = indexedDB.open(dbName);
		request.onerror = (event: Event) => {
			reject(new Error(`failed to open ${dbName} ${event.target}`));
		};
		request.onupgradeneeded = (event) => {
			const db: IDBDatabase = (event.target as IDBRequest).result;
			const objectStore = db.createObjectStore(objectStoreName, { keyPath });

			// データを追加する前に objectStore の作成を完了させるため、
			// transaction oncomplete を使用します。
			objectStore.transaction.oncomplete = () => {
				// 新たに作成した objectStore に値を保存します。
				const customerObjectStore = db
					.transaction(objectStoreName, 'readwrite')
					.objectStore(objectStoreName);
				customerObjectStore.add(obj);
			};
		};
	});
};
