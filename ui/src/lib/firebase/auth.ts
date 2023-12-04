import { getAuth, connectAuthEmulator } from 'firebase/auth';
import { app } from './app';

export const auth = getAuth(app);
connectAuthEmulator(auth, 'http://127.0.0.1:9099');
