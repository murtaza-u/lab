import type { Post } from "../api/post/+server";

export const load = async ({ fetch }) => {
    const response = await fetch("/api/post")
    const posts = await response.json() as Post[]
    return {
        "posts": posts
    }
}