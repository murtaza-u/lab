<script lang="ts">
	import { count } from './store';
</script>

<button on:click={() => count.reset()}>Reset</button>
