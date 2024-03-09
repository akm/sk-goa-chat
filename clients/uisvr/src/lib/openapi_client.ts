// https://github.com/drwpow/openapi-typescript/tree/main/packages/openapi-fetch
import createClient from 'openapi-fetch';

import type { paths } from './openapi';

const client = createClient<paths>();

export const GET = client.GET;
export const POST = client.POST;
export const PUT = client.PUT;
export const DELETE = client.DELETE;
