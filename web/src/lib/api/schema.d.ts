/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
	"/api/files": {
		parameters: {
			query?: never;
			header?: never;
			path?: never;
			cookie?: never;
		};
		/** List all available files */
		get: operations["filesList"];
		put?: never;
		/** Upload a file to Oxidrive */
		post: operations["filesUpload"];
		delete?: never;
		options?: never;
		head?: never;
		patch?: never;
		trace?: never;
	};
	"/api/files/{id}": {
		parameters: {
			query?: never;
			header?: never;
			path?: never;
			cookie?: never;
		};
		get?: never;
		put?: never;
		post?: never;
		delete: operations["fileDelete"];
		options?: never;
		head?: never;
		patch?: never;
		trace?: never;
	};
	"/api/instance": {
		parameters: {
			query?: never;
			header?: never;
			path?: never;
			cookie?: never;
		};
		/** Get the instance status */
		get: operations["instanceStatus"];
		put?: never;
		post?: never;
		delete?: never;
		options?: never;
		head?: never;
		patch?: never;
		trace?: never;
	};
	"/api/instance/setup": {
		parameters: {
			query?: never;
			header?: never;
			path?: never;
			cookie?: never;
		};
		get?: never;
		put?: never;
		/** Setup the instance and create the initial admin user */
		post: operations["instanceSetup"];
		delete?: never;
		options?: never;
		head?: never;
		patch?: never;
		trace?: never;
	};
	"/api/session": {
		parameters: {
			query?: never;
			header?: never;
			path?: never;
			cookie?: never;
		};
		/** Return the user information related to the current session */
		get: operations["authGetSession"];
		put?: never;
		post?: never;
		delete?: never;
		options?: never;
		head?: never;
		patch?: never;
		trace?: never;
	};
	"/api/sessions": {
		parameters: {
			query?: never;
			header?: never;
			path?: never;
			cookie?: never;
		};
		get?: never;
		put?: never;
		/** Create a new session and generate the corresponding token */
		post: operations["authCreateSession"];
		delete?: never;
		options?: never;
		head?: never;
		patch?: never;
		trace?: never;
	};
}
export type webhooks = Record<string, never>;
export interface components {
	schemas: {
		Error: {
			/** @description machine-readable error tag */
			error: string;
			/** @description human readable error message */
			message: string;
		};
		NotFoundError: {
			/** @enum {string} */
			error: "not_found";
			/** @description human readable error message */
			message: string;
		};
		ListInfo: {
			/**
			 * @description number of items in the current slice of the collection
			 * @example 1
			 */
			count: number;
			/**
			 * @description total number of items in the collection
			 * @example 42
			 */
			total: number;
			/**
			 * @description Cursor of the next element, to be used as the `after` parameter in paginated operations
			 * @example 6f6bed9b-320e-4917-a90a-e26e471c4a52
			 */
			next: string | null;
		};
		File: {
			/** Format: uuid */
			id: string;
			/** @enum {string} */
			type: "file" | "folder";
			/** @example image/png */
			contentType: string;
			/**
			 * Format: path
			 * @example path/to/file.txt
			 */
			path: string;
			/** @example file.txt */
			name: string;
			/**
			 * @description Size of the file in bytes
			 * @example 42
			 */
			size: number;
		};
		FileList: components["schemas"]["ListInfo"] & {
			items: components["schemas"]["File"][];
		};
		FileUpload: {
			path: string;
			/** Format: binary */
			file: string;
		};
		FileUploadResponse: {
			ok: boolean;
			id: string;
		};
		InstanceStatus: {
			status: {
				/**
				 * Format: uri
				 * @example https://storage.example.com
				 */
				publicURL: string;
				/** @enum {string} */
				database: "postgres" | "sqlite";
				/** @enum {string} */
				fileStorage: "filesystem" | "s3";
				setupCompleted: boolean;
			};
		};
		InstanceSetupRequest: {
			admin: {
				username: string;
				password: string;
			};
		};
		InstanceSetupResponse: {
			ok: boolean;
		};
		SessionRequest: {
			credentials: components["schemas"]["Credentials"];
		};
		Session: {
			user: components["schemas"]["User"];
			/** Format: date-time */
			expiresAt: string;
		};
		User: {
			/** Format: uuid */
			id: string;
			username: string;
		};
		Credentials: {
			/** @enum {string} */
			kind: "password";
		} & components["schemas"]["PasswordCredentials"];
		PasswordCredentials: {
			/** @enum {string} */
			kind: "password";
			username: string;
			password: string;
		};
	};
	responses: {
		Error: {
			headers: {
				[name: string]: unknown;
			};
			content: {
				"application/json": components["schemas"]["Error"];
			};
		};
		NotFound: {
			headers: {
				[name: string]: unknown;
			};
			content: {
				"application/json": components["schemas"]["NotFoundError"];
			};
		};
		/** @description Unexpected Error */
		InternalError: {
			headers: {
				[name: string]: unknown;
			};
			content: {
				"application/json": components["schemas"]["Error"];
			};
		};
	};
	parameters: {
		/** @description Cursor to fetch the next slice of the collection */
		After: string;
		/** @description Limit the number of items to return to only the first N */
		First: number;
		/** @description Prefix to filter files for. This is matched against the directory the files resides in, not as a generic prefix.
		 *     E.g. a prefix `hello` will match `hello/world.txt` but not `hello/dear/world.txt`.
		 *      */
		FilePrefix: string;
	};
	requestBodies: never;
	headers: never;
	pathItems: never;
}
export type $defs = Record<string, never>;
export interface operations {
	filesList: {
		parameters: {
			query?: {
				/** @description Limit the number of items to return to only the first N */
				first?: components["parameters"]["First"];
				/** @description Cursor to fetch the next slice of the collection */
				after?: components["parameters"]["After"];
				/** @description Prefix to filter files for. This is matched against the directory the files resides in, not as a generic prefix.
				 *     E.g. a prefix `hello` will match `hello/world.txt` but not `hello/dear/world.txt`.
				 *      */
				prefix?: components["parameters"]["FilePrefix"];
			};
			header?: never;
			path?: never;
			cookie?: never;
		};
		requestBody?: never;
		responses: {
			200: {
				headers: {
					[name: string]: unknown;
				};
				content: {
					"application/json": components["schemas"]["FileList"];
				};
			};
			default: components["responses"]["InternalError"];
		};
	};
	filesUpload: {
		parameters: {
			query?: never;
			header?: never;
			path?: never;
			cookie?: never;
		};
		requestBody: {
			content: {
				"multipart/form-data": components["schemas"]["FileUpload"];
			};
		};
		responses: {
			200: {
				headers: {
					[name: string]: unknown;
				};
				content: {
					"application/json": components["schemas"]["FileUploadResponse"];
				};
			};
			400: components["responses"]["Error"];
			default: components["responses"]["InternalError"];
		};
	};
	fileDelete: {
		parameters: {
			query?: never;
			header?: never;
			path: {
				/** @description Id of the file to delete */
				id: string;
			};
			cookie?: never;
		};
		requestBody?: never;
		responses: {
			200: {
				headers: {
					[name: string]: unknown;
				};
				content: {
					"application/json": components["schemas"]["File"];
				};
			};
			/** @description The requested file does not exist or cannot be accessed */
			404: components["responses"]["NotFound"];
			default: components["responses"]["InternalError"];
		};
	};
	instanceStatus: {
		parameters: {
			query?: never;
			header?: never;
			path?: never;
			cookie?: never;
		};
		requestBody?: never;
		responses: {
			200: {
				headers: {
					[name: string]: unknown;
				};
				content: {
					"application/json": components["schemas"]["InstanceStatus"];
				};
			};
			default: components["responses"]["InternalError"];
		};
	};
	instanceSetup: {
		parameters: {
			query?: never;
			header?: never;
			path?: never;
			cookie?: never;
		};
		requestBody: {
			content: {
				"application/json": components["schemas"]["InstanceSetupRequest"];
			};
		};
		responses: {
			200: {
				headers: {
					[name: string]: unknown;
				};
				content: {
					"application/json": components["schemas"]["InstanceSetupResponse"];
				};
			};
			400: components["responses"]["Error"];
			/** @description Returned if the first time setup flow has already been completed */
			409: components["responses"]["Error"];
			default: components["responses"]["InternalError"];
		};
	};
	authGetSession: {
		parameters: {
			query?: never;
			header?: never;
			path?: never;
			cookie?: never;
		};
		requestBody?: never;
		responses: {
			200: {
				headers: {
					[name: string]: unknown;
				};
				content: {
					"application/json": components["schemas"]["User"];
				};
			};
		};
	};
	authCreateSession: {
		parameters: {
			query?: never;
			header?: never;
			path?: never;
			cookie?: never;
		};
		requestBody: {
			content: {
				"application/json": components["schemas"]["SessionRequest"];
			};
		};
		responses: {
			200: {
				headers: {
					"Set-Cookie"?: string;
					[name: string]: unknown;
				};
				content: {
					"application/json": components["schemas"]["Session"];
				};
			};
			/** @description The supplied credentials are invalid */
			401: components["responses"]["Error"];
			default: components["responses"]["InternalError"];
		};
	};
}
