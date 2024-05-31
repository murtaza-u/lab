import type { Config } from 'drizzle-kit'

// fill me
const host = '35.154.95.212'
const user = 'murtaza'
const password =
	'Rz5xXVAHWCCbzQrJxbTrLebdwobcSr756nPQTC3SnErY3CBMKoKcyUwAumA3udLy47Y48h'
const db = 'knowcode'

export default {
	schema: './src/lib/server/schema.ts',
	out: './drizzle',
	driver: 'mysql2',
	dbCredentials: {
		host: host,
		user: user,
		password: password,
		database: db,
		port: 3306
	}
} satisfies Config
