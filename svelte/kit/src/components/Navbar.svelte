<script lang="ts">
	import { page } from '$app/stores'

	$: path = $page.route.id
</script>

<nav>
	<div>
		<h1>Blog</h1>
	</div>
	<div>
		<ul>
			<li><a class:active={path === '/'} href="/">Home</a></li>
			<li><a class:active={path === '/posts'} href="/posts">Posts</a></li>
		</ul>
	</div>
</nav>

<style>
	nav {
		padding: 1rem;
	}

	.active {
		border-bottom: 1px solid;
		border-radius: 0;
	}
</style>
