import { drizzle } from 'drizzle-orm/mysql2'
import mysql from 'mysql2/promise'
import {
	MYSQL_HOST,
	MYSQL_USER,
	MYSQL_PASSWORD,
	MYSQL_DATABASE
} from '$env/static/private'

export const pool = mysql.createPool({
	host: MYSQL_HOST,
	user: MYSQL_USER,
	password: MYSQL_PASSWORD,
	database: MYSQL_DATABASE
})

export const db = drizzle(pool)
