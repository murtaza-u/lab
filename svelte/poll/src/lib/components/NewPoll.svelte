<script lang="ts">
	import { pollStore } from '$lib/store';
	import { toastStore } from '@skeletonlabs/skeleton';
	import type { ToastSettings } from '@skeletonlabs/skeleton';

	const t: ToastSettings = {
		message: 'Added new poll'
	};

	let formData = { question: '', opt1: '', opt2: '' };

	function handler() {
		pollStore.append({
			question: formData.question,
			opt1: formData.opt1,
			opt2: formData.opt2,
			opt1Votes: 0,
			opt2Votes: 0,
			id: 0
		});
		formData = { question: '', opt1: '', opt2: '' };
		toastStore.trigger(t);
	}
</script>

<form class="w-3/5 mx-auto mt-10" on:submit|preventDefault={handler}>
	<label class="label my-5">
		<span>Question</span>
		<input
			class="input px-5 py-2 rounded-md"
			title="Question"
			type="text"
			placeholder="Type question"
			required
			bind:value={formData.question}
		/>
	</label>

	<label class="label my-5">
		<span>Option 1</span>
		<input
			class="input px-5 py-2 rounded-md"
			title="Option 1"
			type="text"
			placeholder="Type first option"
			required
			bind:value={formData.opt1}
		/>
	</label>

	<label class="label my-5">
		<span>Option 2</span>
		<input
			class="input px-5 py-2 rounded-md"
			title="Option 2"
			type="text"
			placeholder="Type second option"
			required
			bind:value={formData.opt2}
		/>
	</label>

	<button type="submit" class="btn variant-filled float-right mt-5 rounded-md">Add</button>
</form>
