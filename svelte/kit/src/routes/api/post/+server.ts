import { error, type RequestHandler } from "@sveltejs/kit";

export type Post = {
    id: number
    date: Date
    title: string
    body: string
    slug: string
}

const posts: Post[] = [
    {
        id: 1,
        date: new Date(Date.now()),
        title: "Learn svelte",
        body: "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Labore fuga magni corporis molestiae quas amet. Illo laboriosam iusto inventore facere dicta obcaecati fugiat quos maxime dolorum magnam eveniet, laudantium animi!",
        slug: "learn-svelte",
    },
    {
        id: 2,
        date: new Date(Date.now()),
        title: "Learn react",
        body: "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Labore fuga magni corporis molestiae quas amet. Illo laboriosam iusto inventore facere dicta obcaecati fugiat quos maxime dolorum magnam eveniet, laudantium animi!",
        slug: "learn-react",
    },
    {
        id: 3,
        date: new Date(Date.now()),
        title: "Learn angular",
        body: "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Labore fuga magni corporis molestiae quas amet. Illo laboriosam iusto inventore facere dicta obcaecati fugiat quos maxime dolorum magnam eveniet, laudantium animi!",
        slug: "learn-angular",
    },
    {
        id: 4,
        date: new Date(Date.now()),
        title: "Learn vue",
        body: "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Labore fuga magni corporis molestiae quas amet. Illo laboriosam iusto inventore facere dicta obcaecati fugiat quos maxime dolorum magnam eveniet, laudantium animi!",
        slug: "learn-vue",
    },
    {
        id: 5,
        date: new Date(Date.now()),
        title: "Learn Go",
        body: "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Labore fuga magni corporis molestiae quas amet. Illo laboriosam iusto inventore facere dicta obcaecati fugiat quos maxime dolorum magnam eveniet, laudantium animi!",
        slug: "learn-vue",
    },
    {
        id: 6,
        date: new Date(Date.now()),
        title: "Learn vue",
        body: "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Labore fuga magni corporis molestiae quas amet. Illo laboriosam iusto inventore facere dicta obcaecati fugiat quos maxime dolorum magnam eveniet, laudantium animi!",
        slug: "learn-vue",
    },
    {
        id: 7,
        date: new Date(Date.now()),
        title: "Learn vue",
        body: "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Labore fuga magni corporis molestiae quas amet. Illo laboriosam iusto inventore facere dicta obcaecati fugiat quos maxime dolorum magnam eveniet, laudantium animi!",
        slug: "learn-vue",
    },
    {
        id: 8,
        date: new Date(Date.now()),
        title: "Learn vue",
        body: "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Labore fuga magni corporis molestiae quas amet. Illo laboriosam iusto inventore facere dicta obcaecati fugiat quos maxime dolorum magnam eveniet, laudantium animi!",
        slug: "learn-vue",
    },
]

export const GET: RequestHandler = ({ url }): Response => {
    const limit = Number(url.searchParams.get("limit") ?? 30)
    const slug = url.searchParams.get("slug")
    const idx = limit > posts.length ? posts.length : limit

    if (slug) {
        const post = posts.find(post => post.slug === slug)
        if (!post) throw error(404, `post ${slug} does not exists`)
        return new Response(JSON.stringify(post), {
            headers: {
                "Content-Type": "application/json",
            },
        })
    }

    return new Response(JSON.stringify(posts.slice(0, idx)), {
        headers: {
            "Content-Type": "application/json",
        },
    })
}