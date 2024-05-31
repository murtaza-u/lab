import { fail, redirect } from '@sveltejs/kit'

export const actions = {
    default: async ({ url, locals: { supabase } }) => {
        const { data, error } = await supabase.auth.signInWithOAuth({
            provider: 'google',
            options: {
                redirectTo: `${url.origin}/auth/callback`,
                queryParams: {
                    access_type: 'offline',
                    prompt: 'consent',
                },
            }
        })

        if (error) {
            return fail(500, {
                message: error.message
            })
        }

        throw redirect(303, data.url)
    },
}