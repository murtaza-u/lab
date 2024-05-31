import { lucia } from 'lucia'
import { sveltekit } from 'lucia/middleware'
import { dev } from '$app/environment'
import { mysql2 } from '@lucia-auth/adapter-mysql'
import { pool } from './db'

export const auth = lucia({
	env: dev ? 'DEV' : 'PROD',
	middleware: sveltekit(),
	adapter: mysql2(pool, {
		user: 'auth_user',
		key: 'auth_key',
		session: 'auth_session'
	}),
	getUserAttributes: (data) => {
		return {
			username: data.username
		}
	}
})

export type Auth = typeof auth
