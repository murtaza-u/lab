import type { Post } from "../../api/post/+server";

export const load = async ({ fetch, params }) => {
    const slug = params.slug
    const response = await fetch(`/api/post?slug=${slug}`)
    const post = await response.json() as Post
    return post
}