import { auth } from '$lib/server/lucia'
import { LuciaError } from 'lucia'
import { fail, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

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

		try {
			// find user by key and validate password
			const key = await auth.useKey(
				'username',
				username.toLowerCase(),
				password
			)
			const session = await auth.createSession({
				userId: key.userId,
				attributes: {}
			})
			await auth.deleteDeadUserSessions(key.userId)
			locals.auth.setSession(session)
		} catch (e) {
			if (
				e instanceof LuciaError &&
				(e.message === 'AUTH_INVALID_KEY_ID' ||
					e.message === 'AUTH_INVALID_PASSWORD')
			) {
				// user does not exist
				// or invalid password
				return fail(400, {
					ctx: 'global',
					error: 'Incorrect username or password'
				})
			}
			return fail(500, {
				ctx: 'global',
				error: 'An unknown error occurred'
			})
		}
		throw redirect(302, '/')
	}
}
