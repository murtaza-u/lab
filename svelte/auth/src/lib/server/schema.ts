import { bigint, mysqlTable, varchar } from 'drizzle-orm/mysql-core'

export const userTable = mysqlTable('auth_user', {
	id: varchar('id', { length: 15 }).primaryKey().notNull(),
	username: varchar('username', { length: 50 }).notNull()
})

export const keyTable = mysqlTable('auth_key', {
	id: varchar('id', { length: 255 }).primaryKey().notNull(),
	userId: varchar('user_id', { length: 15 })
		.notNull()
		.references(() => userTable.id),
	hashedPassword: varchar('hashed_password', { length: 255 })
})

export const sessionTable = mysqlTable('auth_session', {
	id: varchar('id', { length: 127 }).primaryKey().notNull(),
	userId: varchar('user_id', { length: 15 })
		.notNull()
		.references(() => userTable.id),
	activeExpires: bigint('active_expires', { mode: 'bigint' }).notNull(),
	idelExpires: bigint('idle_expires', { mode: 'bigint' }).notNull()
})
