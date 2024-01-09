export type ChatMessage = {
	id: bigint;
	channelId: bigint;
	userId?: bigint;
	userName: string;
	content: string;
	createdAt: string;
	updatedAt: string;
};
