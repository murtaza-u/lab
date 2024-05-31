import { localStorageStore } from '@skeletonlabs/skeleton';

type Poll = {
    id: number
    question: string
    opt1: string
    opt2: string
    opt1Votes: number
    opt2Votes: number
}

function newPollStore() {
    const polls: Poll[] = []
    const { subscribe, update } = localStorageStore("polls", polls)

    return {
        subscribe,
        append: (poll: Poll): void => {
            update(list => [...list, { ...poll, id: list.length + 1 }])
        },
        delete: (id: number): void => {
            update(list => list.filter(poll => poll.id !== id))
        },
        incOpt1Vote: (id: number): void => {
            update(list => list.map(poll => {
                if (poll.id !== id) return poll
                poll.opt1Votes++
                return poll
            }))
        },
        incOpt2Vote: (id: number): void => {
            update(list => list.map(poll => {
                if (poll.id !== id) return poll
                poll.opt2Votes++
                return poll
            }))
        },
    }
}

export const pollStore = newPollStore();