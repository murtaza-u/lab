import { writable } from "svelte/store";

function createCount() {
    const { subscribe, set, update } = writable(0)

    return {
        subscribe,
        set,
        reset: () => set(0),
        increment: () => update(c => c + 1),
        decrement: () => update(c => c - 1),
    }
}

export const count = createCount()