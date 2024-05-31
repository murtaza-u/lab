<script lang="ts">
	import NewPoll from '$lib/components/NewPoll.svelte';
	import Vote from '$lib/components/Vote.svelte';
	import { TabGroup, Tab } from '@skeletonlabs/skeleton';

	let tabSet: number = 0;
</script>

<TabGroup justify="justify-center">
	<Tab value={0} bind:group={tabSet} name="tab0">Vote</Tab>
	<Tab value={1} bind:group={tabSet} name="tab1">New Poll</Tab>
	<!-- Tab Panels --->
	<svelte:fragment slot="panel">
		{#if tabSet === 0}
			<Vote />
		{:else if tabSet === 1}
			<NewPoll />
		{/if}
	</svelte:fragment>
</TabGroup>
