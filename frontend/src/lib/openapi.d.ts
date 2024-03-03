/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export interface paths {
  "/api/channels": {
    /** list channels */
    get: operations["channels#list"];
    /** create channels */
    post: operations["channels#create"];
  };
  "/api/channels/{id}": {
    /** show channels */
    get: operations["channels#show"];
    /** update channels */
    put: operations["channels#update"];
    /** delete channels */
    delete: operations["channels#delete"];
  };
  "/api/chat_messages": {
    /** list chat_messages */
    get: operations["chat_messages#list"];
    /** create chat_messages */
    post: operations["chat_messages#create"];
  };
  "/api/chat_messages/{id}": {
    /** show chat_messages */
    get: operations["chat_messages#show"];
    /** update chat_messages */
    put: operations["chat_messages#update"];
    /** delete chat_messages */
    delete: operations["chat_messages#delete"];
  };
  "/api/users": {
    /** list users */
    get: operations["users#list"];
    /** create users */
    post: operations["users#create"];
  };
  "/ws/notifications/subscribe": {
    /**
     * subscribe notifications
     * @description Subscribe to events sent such new chat messages.
     */
    get: operations["notifications#subscribe"];
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    /**
     * @example {
     *   "created_at": "2006-01-02T15:04:05Z07:00",
     *   "id": 15377353183551013000,
     *   "name": "Eum consequatur commodi.",
     *   "updated_at": "2006-01-02T15:04:05Z07:00"
     * }
     */
    Channel: {
      /**
       * Format: date-time
       * @description CreatedAt
       * @example 2006-01-02T15:04:05Z07:00
       */
      created_at: string;
      /**
       * @description ID
       * @example 1806459186028017400
       */
      id: number;
      /**
       * @description Name
       * @example Omnis quisquam rem qui officia.
       */
      name: string;
      /**
       * Format: date-time
       * @description UpdatedAt
       * @example 2006-01-02T15:04:05Z07:00
       */
      updated_at: string;
    };
    /**
     * @example {
     *   "items": [
     *     {
     *       "created_at": "2006-01-02T15:04:05Z07:00",
     *       "id": 14308298621592037000,
     *       "name": "Voluptas error.",
     *       "updated_at": "2006-01-02T15:04:05Z07:00"
     *     },
     *     {
     *       "created_at": "2006-01-02T15:04:05Z07:00",
     *       "id": 14308298621592037000,
     *       "name": "Voluptas error.",
     *       "updated_at": "2006-01-02T15:04:05Z07:00"
     *     }
     *   ],
     *   "offset": 0,
     *   "total": 160
     * }
     */
    ChannelList: {
      items: components["schemas"]["ChannelListItemCollection"];
      /**
       * @description Offset
       * @example 0
       */
      offset: number;
      /**
       * @description Total number of items
       * @example 160
       */
      total: number;
    };
    /**
     * @example {
     *   "created_at": "2006-01-02T15:04:05Z07:00",
     *   "id": 6697053573302417000,
     *   "name": "Eaque reprehenderit.",
     *   "updated_at": "2006-01-02T15:04:05Z07:00"
     * }
     */
    ChannelListItem: {
      /**
       * Format: date-time
       * @description CreatedAt
       * @example 2006-01-02T15:04:05Z07:00
       */
      created_at: string;
      /**
       * @description ID
       * @example 1718773759709988900
       */
      id: number;
      /**
       * @description Name
       * @example Aut quasi ut in nihil maiores.
       */
      name: string;
      /**
       * Format: date-time
       * @description UpdatedAt
       * @example 2006-01-02T15:04:05Z07:00
       */
      updated_at: string;
    };
    /**
     * @example [
     *   {
     *     "created_at": "2006-01-02T15:04:05Z07:00",
     *     "id": 14308298621592037000,
     *     "name": "Voluptas error.",
     *     "updated_at": "2006-01-02T15:04:05Z07:00"
     *   },
     *   {
     *     "created_at": "2006-01-02T15:04:05Z07:00",
     *     "id": 14308298621592037000,
     *     "name": "Voluptas error.",
     *     "updated_at": "2006-01-02T15:04:05Z07:00"
     *   },
     *   {
     *     "created_at": "2006-01-02T15:04:05Z07:00",
     *     "id": 14308298621592037000,
     *     "name": "Voluptas error.",
     *     "updated_at": "2006-01-02T15:04:05Z07:00"
     *   }
     * ]
     */
    ChannelListItemCollection: components["schemas"]["ChannelListItem"][];
    /**
     * @example {
     *   "channel_id": 1810376289941694500,
     *   "content": "Iusto ut.",
     *   "created_at": "2006-01-02T15:04:05Z07:00",
     *   "id": 2549351506480365600,
     *   "updated_at": "2006-01-02T15:04:05Z07:00",
     *   "user_id": 14122095824598073000,
     *   "user_name": "Consequatur similique autem voluptatem aut nihil deserunt."
     * }
     */
    ChatMessage: {
      /**
       * @description Channel ID
       * @example 16702029836893956000
       */
      channel_id: number;
      /**
       * @description Content
       * @example Quisquam et fuga omnis maxime.
       */
      content: string;
      /**
       * Format: date-time
       * @description CreatedAt
       * @example 2006-01-02T15:04:05Z07:00
       */
      created_at: string;
      /**
       * @description ID
       * @example 15924806015104549000
       */
      id: number;
      /**
       * Format: date-time
       * @description UpdatedAt
       * @example 2006-01-02T15:04:05Z07:00
       */
      updated_at: string;
      /**
       * @description User ID
       * @example 1719468543456683800
       */
      user_id?: number;
      /**
       * @description User Name
       * @example Aut earum voluptas sit dolores harum quae.
       */
      user_name: string;
    };
    /**
     * @example {
     *   "items": [
     *     {
     *       "channel_id": 13959009638458442000,
     *       "content": "Dolore ut est.",
     *       "created_at": "2006-01-02T15:04:05Z07:00",
     *       "id": 15936474696563802000,
     *       "updated_at": "2006-01-02T15:04:05Z07:00",
     *       "user_id": 16573384439575351000,
     *       "user_name": "Iusto sed fugit beatae nihil."
     *     },
     *     {
     *       "channel_id": 13959009638458442000,
     *       "content": "Dolore ut est.",
     *       "created_at": "2006-01-02T15:04:05Z07:00",
     *       "id": 15936474696563802000,
     *       "updated_at": "2006-01-02T15:04:05Z07:00",
     *       "user_id": 16573384439575351000,
     *       "user_name": "Iusto sed fugit beatae nihil."
     *     },
     *     {
     *       "channel_id": 13959009638458442000,
     *       "content": "Dolore ut est.",
     *       "created_at": "2006-01-02T15:04:05Z07:00",
     *       "id": 15936474696563802000,
     *       "updated_at": "2006-01-02T15:04:05Z07:00",
     *       "user_id": 16573384439575351000,
     *       "user_name": "Iusto sed fugit beatae nihil."
     *     }
     *   ]
     * }
     */
    ChatMessageList: {
      items: components["schemas"]["ChatMessageListItemCollection"];
    };
    /**
     * @example {
     *   "channel_id": 8314486435182394000,
     *   "content": "Voluptate sint tempora.",
     *   "created_at": "2006-01-02T15:04:05Z07:00",
     *   "id": 6451461773391010000,
     *   "updated_at": "2006-01-02T15:04:05Z07:00",
     *   "user_id": 6683625257867510000,
     *   "user_name": "Nulla sed."
     * }
     */
    ChatMessageListItem: {
      /**
       * @description Channel ID
       * @example 2794964574277862400
       */
      channel_id: number;
      /**
       * @description Content
       * @example Atque quo enim unde.
       */
      content: string;
      /**
       * Format: date-time
       * @description CreatedAt
       * @example 2006-01-02T15:04:05Z07:00
       */
      created_at: string;
      /**
       * @description ID
       * @example 17851874401144338000
       */
      id: number;
      /**
       * Format: date-time
       * @description UpdatedAt
       * @example 2006-01-02T15:04:05Z07:00
       */
      updated_at: string;
      /**
       * @description User ID
       * @example 16406711370279236000
       */
      user_id?: number;
      /**
       * @description User Name
       * @example Impedit officiis omnis quia animi odit provident.
       */
      user_name: string;
    };
    /**
     * @example [
     *   {
     *     "channel_id": 13959009638458442000,
     *     "content": "Dolore ut est.",
     *     "created_at": "2006-01-02T15:04:05Z07:00",
     *     "id": 15936474696563802000,
     *     "updated_at": "2006-01-02T15:04:05Z07:00",
     *     "user_id": 16573384439575351000,
     *     "user_name": "Iusto sed fugit beatae nihil."
     *   },
     *   {
     *     "channel_id": 13959009638458442000,
     *     "content": "Dolore ut est.",
     *     "created_at": "2006-01-02T15:04:05Z07:00",
     *     "id": 15936474696563802000,
     *     "updated_at": "2006-01-02T15:04:05Z07:00",
     *     "user_id": 16573384439575351000,
     *     "user_name": "Iusto sed fugit beatae nihil."
     *   },
     *   {
     *     "channel_id": 13959009638458442000,
     *     "content": "Dolore ut est.",
     *     "created_at": "2006-01-02T15:04:05Z07:00",
     *     "id": 15936474696563802000,
     *     "updated_at": "2006-01-02T15:04:05Z07:00",
     *     "user_id": 16573384439575351000,
     *     "user_name": "Iusto sed fugit beatae nihil."
     *   }
     * ]
     */
    ChatMessageListItemCollection: components["schemas"]["ChatMessageListItem"][];
    /**
     * @example {
     *   "name": "Quis ut error quia sit omnis."
     * }
     */
    CreateRequestBody: {
      /**
       * @description Name
       * @example Quo eum in aperiam quia nisi suscipit.
       */
      name: string;
    };
    /**
     * @example {
     *   "channel_id": 17976362873503885000,
     *   "content": "Dolores dolorum tempore facere occaecati sunt rerum."
     * }
     */
    CreateRequestBody2: {
      /**
       * @description Channel ID
       * @example 5281351726259432000
       */
      channel_id: number;
      /**
       * @description Content
       * @example Rerum pariatur nesciunt.
       */
      content: string;
    };
    /**
     * @example {
     *   "email": "Et quaerat aliquid enim asperiores est deserunt.",
     *   "name": "Fugiat voluptas."
     * }
     */
    CreateRequestBody3: {
      /**
       * @description Email
       * @example Ipsum et quia est fugiat.
       */
      email: string;
      /**
       * @description Name
       * @example Minima optio itaque.
       */
      name: string;
    };
    /**
     * @example {
     *   "fault": false,
     *   "id": "123abc",
     *   "message": "parameter 'p' must be an integer",
     *   "name": "bad_request",
     *   "temporary": true,
     *   "timeout": true
     * }
     */
    Error: {
      /**
       * @description Is the error a server-side fault?
       * @example false
       */
      fault: boolean;
      /**
       * @description ID is a unique identifier for this particular occurrence of the problem.
       * @example 123abc
       */
      id: string;
      /**
       * @description Message is a human-readable explanation specific to this occurrence of the problem.
       * @example parameter 'p' must be an integer
       */
      message: string;
      /**
       * @description Name is the name of this class of errors.
       * @example bad_request
       */
      name: string;
      /**
       * @description Is the error temporary?
       * @example false
       */
      temporary: boolean;
      /**
       * @description Is the error a timeout?
       * @example true
       */
      timeout: boolean;
    };
    /**
     * @example {
     *   "channel_ids": [
     *     16695547136315716000,
     *     17426452557373782000,
     *     13700264418383905000
     *   ]
     * }
     */
    NotificationEvent: {
      /**
       * @description IDs of channels which got new messages
       * @example [
       *   11663134532334922000,
       *   1125975758058537000
       * ]
       */
      channel_ids: number[];
    };
    /**
     * @example {
     *   "content": "Qui non quis."
     * }
     */
    UpdateRequestBody: {
      /**
       * @description Content
       * @example Facilis inventore ut mollitia.
       */
      content: string;
    };
    /**
     * @example {
     *   "created_at": "2006-01-02T15:04:05Z07:00",
     *   "id": 17583520077888694000,
     *   "name": "Odit qui non.",
     *   "updated_at": "2006-01-02T15:04:05Z07:00"
     * }
     */
    User: {
      /**
       * Format: date-time
       * @description CreatedAt
       * @example 2006-01-02T15:04:05Z07:00
       */
      created_at: string;
      /**
       * @description ID
       * @example 5593925075509685000
       */
      id: number;
      /**
       * @description Name
       * @example Reiciendis est.
       */
      name: string;
      /**
       * Format: date-time
       * @description UpdatedAt
       * @example 2006-01-02T15:04:05Z07:00
       */
      updated_at: string;
    };
    /**
     * @example {
     *   "items": [
     *     {
     *       "id": 12732113688530887000,
     *       "name": "Est voluptate et numquam quia perferendis."
     *     },
     *     {
     *       "id": 12732113688530887000,
     *       "name": "Est voluptate et numquam quia perferendis."
     *     },
     *     {
     *       "id": 12732113688530887000,
     *       "name": "Est voluptate et numquam quia perferendis."
     *     }
     *   ],
     *   "offset": 0,
     *   "total": 160
     * }
     */
    UserList: {
      items: components["schemas"]["UserListItemCollection"];
      /**
       * @description Offset
       * @example 0
       */
      offset: number;
      /**
       * @description Total number of items
       * @example 160
       */
      total: number;
    };
    /**
     * @example {
     *   "id": 13456660312305150000,
     *   "name": "Quis architecto reprehenderit quae repellendus ut."
     * }
     */
    UserListItem: {
      /**
       * @description ID
       * @example 15511322544835242000
       */
      id: number;
      /**
       * @description Name
       * @example Ipsum cum accusantium.
       */
      name: string;
    };
    /**
     * @example [
     *   {
     *     "id": 12732113688530887000,
     *     "name": "Est voluptate et numquam quia perferendis."
     *   },
     *   {
     *     "id": 12732113688530887000,
     *     "name": "Est voluptate et numquam quia perferendis."
     *   },
     *   {
     *     "id": 12732113688530887000,
     *     "name": "Est voluptate et numquam quia perferendis."
     *   }
     * ]
     */
    UserListItemCollection: components["schemas"]["UserListItem"][];
  };
  responses: never;
  parameters: never;
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export interface operations {

  /** list channels */
  "channels#list": {
    parameters: {
      header: {
        /**
         * @description X-ID-TOKEN
         * @example abcdef12345
         */
        "X-ID-TOKEN": string;
      };
    };
    responses: {
      /** @description OK response. */
      200: {
        content: {
          "application/json": components["schemas"]["ChannelList"];
        };
      };
      /** @description unauthenticated: Unauthorized response. */
      401: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
    };
  };
  /** create channels */
  "channels#create": {
    parameters: {
      header: {
        /**
         * @description X-ID-TOKEN
         * @example abcdef12345
         */
        "X-ID-TOKEN": string;
      };
    };
    requestBody: {
      content: {
        /**
         * @example {
         *   "name": "Dolorem optio consequatur est eligendi."
         * }
         */
        "application/json": components["schemas"]["CreateRequestBody"];
      };
    };
    responses: {
      /** @description Created response. */
      201: {
        content: {
          "application/json": components["schemas"]["Channel"];
        };
      };
      /** @description invalid_payload: Bad Request response. */
      400: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
      /** @description unauthenticated: Unauthorized response. */
      401: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
    };
  };
  /** show channels */
  "channels#show": {
    parameters: {
      header: {
        /**
         * @description X-ID-TOKEN
         * @example abcdef12345
         */
        "X-ID-TOKEN": string;
      };
      path: {
        /**
         * @description ID
         * @example 2135540402709783000
         */
        id: number;
      };
    };
    responses: {
      /** @description OK response. */
      200: {
        content: {
          "application/json": components["schemas"]["Channel"];
        };
      };
      /** @description unauthenticated: Unauthorized response. */
      401: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
      /** @description not_found: Not Found response. */
      404: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
    };
  };
  /** update channels */
  "channels#update": {
    parameters: {
      header: {
        /**
         * @description X-ID-TOKEN
         * @example abcdef12345
         */
        "X-ID-TOKEN": string;
      };
      path: {
        /**
         * @description ID
         * @example 57829092557472510
         */
        id: number;
      };
    };
    requestBody: {
      content: {
        /**
         * @example {
         *   "name": "Animi ut aut totam."
         * }
         */
        "application/json": components["schemas"]["CreateRequestBody"];
      };
    };
    responses: {
      /** @description OK response. */
      200: {
        content: {
          "application/json": components["schemas"]["Channel"];
        };
      };
      /** @description invalid_payload: Bad Request response. */
      400: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
      /** @description unauthenticated: Unauthorized response. */
      401: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
      /** @description not_found: Not Found response. */
      404: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
    };
  };
  /** delete channels */
  "channels#delete": {
    parameters: {
      header: {
        /**
         * @description X-ID-TOKEN
         * @example abcdef12345
         */
        "X-ID-TOKEN": string;
      };
      path: {
        /**
         * @description ID
         * @example 3345418820542876000
         */
        id: number;
      };
    };
    responses: {
      /** @description OK response. */
      200: {
        content: {
          "application/json": components["schemas"]["Channel"];
        };
      };
      /** @description unauthenticated: Unauthorized response. */
      401: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
      /** @description not_found: Not Found response. */
      404: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
    };
  };
  /** list chat_messages */
  "chat_messages#list": {
    parameters: {
      query: {
        /**
         * @description Limit
         * @example 1756608611783287600
         */
        limit: number;
        /**
         * @description Channel ID
         * @example 5335993148139461000
         */
        channel_id?: number;
        /**
         * @description ChatMessage ID for query to get messages after this
         * @example 8766397724585147000
         */
        after?: number;
        /**
         * @description ChatMessage ID for query to get messages before this
         * @example 3428741854183539000
         */
        before?: number;
      };
      header: {
        /**
         * @description X-ID-TOKEN
         * @example abcdef12345
         */
        "X-ID-TOKEN": string;
      };
    };
    responses: {
      /** @description OK response. */
      200: {
        content: {
          "application/json": components["schemas"]["ChatMessageList"];
        };
      };
      /** @description unauthenticated: Unauthorized response. */
      401: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
    };
  };
  /** create chat_messages */
  "chat_messages#create": {
    parameters: {
      header: {
        /**
         * @description X-ID-TOKEN
         * @example abcdef12345
         */
        "X-ID-TOKEN": string;
      };
    };
    requestBody: {
      content: {
        /**
         * @example {
         *   "channel_id": 3412152620045203500,
         *   "content": "Non accusantium."
         * }
         */
        "application/json": components["schemas"]["CreateRequestBody2"];
      };
    };
    responses: {
      /** @description Created response. */
      201: {
        content: {
          "application/json": components["schemas"]["ChatMessage"];
        };
      };
      /** @description invalid_payload: Bad Request response. */
      400: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
      /** @description unauthenticated: Unauthorized response. */
      401: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
    };
  };
  /** show chat_messages */
  "chat_messages#show": {
    parameters: {
      header: {
        /**
         * @description X-ID-TOKEN
         * @example abcdef12345
         */
        "X-ID-TOKEN": string;
      };
      path: {
        /**
         * @description ID
         * @example 7640296165610348000
         */
        id: number;
      };
    };
    responses: {
      /** @description OK response. */
      200: {
        content: {
          "application/json": components["schemas"]["ChatMessage"];
        };
      };
      /** @description unauthenticated: Unauthorized response. */
      401: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
      /** @description not_found: Not Found response. */
      404: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
    };
  };
  /** update chat_messages */
  "chat_messages#update": {
    parameters: {
      header: {
        /**
         * @description X-ID-TOKEN
         * @example abcdef12345
         */
        "X-ID-TOKEN": string;
      };
      path: {
        /**
         * @description ID
         * @example 2596968404017913000
         */
        id: number;
      };
    };
    requestBody: {
      content: {
        /**
         * @example {
         *   "content": "Qui similique dolor."
         * }
         */
        "application/json": components["schemas"]["UpdateRequestBody"];
      };
    };
    responses: {
      /** @description OK response. */
      200: {
        content: {
          "application/json": components["schemas"]["ChatMessage"];
        };
      };
      /** @description invalid_payload: Bad Request response. */
      400: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
      /** @description unauthenticated: Unauthorized response. */
      401: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
      /** @description not_found: Not Found response. */
      404: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
    };
  };
  /** delete chat_messages */
  "chat_messages#delete": {
    parameters: {
      header: {
        /**
         * @description X-ID-TOKEN
         * @example abcdef12345
         */
        "X-ID-TOKEN": string;
      };
      path: {
        /**
         * @description ID
         * @example 13200870719650333000
         */
        id: number;
      };
    };
    responses: {
      /** @description OK response. */
      200: {
        content: {
          "application/json": components["schemas"]["ChatMessage"];
        };
      };
      /** @description unauthenticated: Unauthorized response. */
      401: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
      /** @description not_found: Not Found response. */
      404: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
    };
  };
  /** list users */
  "users#list": {
    responses: {
      /** @description OK response. */
      200: {
        content: {
          "application/json": components["schemas"]["UserList"];
        };
      };
    };
  };
  /** create users */
  "users#create": {
    requestBody: {
      content: {
        /**
         * @example {
         *   "email": "Vero in dolore odio quasi.",
         *   "name": "Quo minus quas voluptatibus consequatur nemo."
         * }
         */
        "application/json": components["schemas"]["CreateRequestBody3"];
      };
    };
    responses: {
      /** @description Created response. */
      201: {
        content: {
          "application/json": components["schemas"]["User"];
        };
      };
      /** @description invalid_payload: Bad Request response. */
      400: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
    };
  };
  /**
   * subscribe notifications
   * @description Subscribe to events sent such new chat messages.
   */
  "notifications#subscribe": {
    parameters: {
      header: {
        /**
         * @description X-ID-TOKEN
         * @example abcdef12345
         */
        "X-ID-TOKEN": string;
      };
    };
    responses: {
      /** @description Switching Protocols response. */
      101: {
        content: {
          "application/json": components["schemas"]["NotificationEvent"];
        };
      };
      /** @description unauthenticated: Unauthorized response. */
      401: {
        content: {
          "application/vnd.goa.error": components["schemas"]["Error"];
        };
      };
    };
  };
}
