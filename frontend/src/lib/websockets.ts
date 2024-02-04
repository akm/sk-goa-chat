import 'websocket-polyfill';

let _notificationsSocket: WebSocket;

export const notificationsSocket = (): WebSocket => {
	if (_notificationsSocket) return _notificationsSocket;

	const apisvrOrigin = import.meta.env.VITE_UISVR_ORIGIN as string;
	const wsOrigin = apisvrOrigin.includes('https://')
		? apisvrOrigin.replace('https://', 'wss://')
		: apisvrOrigin.replace('http://', 'ws://');
	const wsUrl = `${wsOrigin}/ws/notifications/subscribe`;

	// const websocket = io(wsUrl);
	const ws = new WebSocket(wsUrl);
	ws.addEventListener('open', (event) => {
		console.log('WebSocket open', event);
	});
	ws.addEventListener('close', (event) => {
		console.log('WebSocket close', event);
	});
	ws.addEventListener('error', (event) => {
		console.log('WebSocket error', event);
	});

	_notificationsSocket = ws;
	return _notificationsSocket;
};

export const closeWebSockets = (): void => {
	if (!_notificationsSocket) return;
	_notificationsSocket.close(1000);
};
