<script lang="ts">
	import { pollStore } from '$lib/store';
	import { fade, scale } from 'svelte/transition';
	import { flip } from 'svelte/animate';
	import ProgressBar from './ProgressBar.svelte';
	import DeleteBinLine from 'svelte-remixicon/lib/icons/DeleteBinLine.svelte';

	function incOpt1(id: number) {
		pollStore.incOpt1Vote(id);
	}

	function incOpt2(id: number) {
		pollStore.incOpt2Vote(id);
	}

	function remove(id: number) {
		pollStore.delete(id);
	}

	function getText(opt: string, votes: number): string {
		const post = votes > 1 ? 'votes' : 'vote';
		return `${opt} (${votes} ${post})`;
	}
</script>

<div class="flex flex-wrap justify-evenly my-10">
	{#each $pollStore as poll (poll.id)}
		<div
			in:fade
			out:scale|local
			animate:flip={{ duration: 300 }}
			class="card lg:w-1/3 w-4/5 p-5 m-2"
		>
			<DeleteBinLine
				color="red"
				size="1.5rem"
				class="cursor-pointer float-right"
				on:click={() => remove(poll.id)}
			/>
			<h3 class="card-header">{poll.question}</h3>
			<section>
				<div class="my-5 cursor-pointer" on:mousedown={() => incOpt1(poll.id)}>
					<ProgressBar
						text={getText(poll.opt1, poll.opt1Votes)}
						value={poll.opt1Votes}
						maxValue={poll.opt1Votes > poll.opt2Votes ? poll.opt1Votes : poll.opt2Votes}
						color="pink"
					/>
				</div>

				<div class="my-5 cursor-pointer" on:mousedown={() => incOpt2(poll.id)}>
					<ProgressBar
						text={getText(poll.opt2, poll.opt2Votes)}
						value={poll.opt2Votes}
						maxValue={poll.opt1Votes > poll.opt2Votes ? poll.opt1Votes : poll.opt2Votes}
						color="lightgreen"
					/>
				</div>
			</section>
		</div>
	{/each}
</div>
