import { initializeApp } from 'firebase/app';
import type { FirebaseError } from 'firebase/app';

// ui/src/lib/firebase/firebaseconfig.ts は Firebase のコンソールで
// プロジェクトの設定 > 全般 > マイアプリ > SDK の設定と構成 の Config を選択して
// 表示されるコードをコピーして作成してください。
import { firebaseConfig } from './firebaseconfig';

export const app = initializeApp(firebaseConfig);

export const isFirebaseError = (error: unknown): error is FirebaseError => {
	return error instanceof Error && 'code' in error && 'message' in error;
};

export type { FirebaseError } from 'firebase/app';
