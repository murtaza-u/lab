import { auth } from '$lib/server/lucia'
import { fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'
import { db } from '$lib/server/db'
import { userTable } from '$lib/server/schema'
import { eq } from 'drizzle-orm'

export const load: PageServerLoad = async ({ locals }) => {
	const session = await locals.auth.validate()
	if (session) throw redirect(302, '/')
	return {}
}

export const actions: Actions = {
	default: async ({ request, locals }) => {
		const formData = await request.formData()
		const username = formData.get('username')
		const password = formData.get('password')
		const confirmPassword = formData.get('confirm_password')

		if (!username || typeof username !== 'string') {
			return fail(400, {
				ctx: 'username',
				error: 'Username cannot be blank'
			})
		}
		if (!password || typeof password !== 'string') {
			return fail(400, {
				ctx: 'password',
				error: 'Password cannot be blank'
			})
		}
		if (!confirmPassword) {
			return fail(400, {
				ctx: 'confirm_password',
				error: 'Confirm password cannot be blank'
			})
		}
		if (password.toString() !== confirmPassword.toString()) {
			return fail(400, {
				ctx: 'confirm_password',
				error: 'Does not match password'
			})
		}

		try {
			const res = await db
				.select({ id: userTable.id })
				.from(userTable)
				.where(eq(userTable.username, username.toString()))

			if (res.length !== 0) {
				return fail(400, {
					ctx: 'username',
					error: 'Username already exists'
				})
			}
		} catch (e) {
			return fail(500, {
				ctx: 'global',
				error: 'An unknown error occurred'
			})
		}

		try {
			const user = await auth.createUser({
				key: {
					providerId: 'username',
					providerUserId: username.toString().toLowerCase(),
					password
				},
				attributes: {
					username
				}
			})
			const session = await auth.createSession({
				userId: user.userId,
				attributes: {}
			})
			locals.auth.setSession(session)
		} catch (e) {
			return fail(500, {
				ctx: 'global',
				error: 'An unknown error occurred'
			})
		}
		throw redirect(302, '/')
	}
}
