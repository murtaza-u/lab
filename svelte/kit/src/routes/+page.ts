import type { Post } from "./api/post/+server";

export const load = async ({ fetch }) => {
    const response = await fetch("/api/post?limit=5")
    const posts = await response.json() as Post[]
    return {
        "posts": posts
    }
}