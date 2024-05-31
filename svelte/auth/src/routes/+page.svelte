<script lang="ts">
	import { enhance } from '$app/forms'
	import type { PageServerData } from './$types'

	export let data: PageServerData
</script>

<header>
	<nav class="navbar bg-base-100">
		<div class="flex-1">
			<a href="/" class="btn btn-ghost text-xl hover:bg-base-100">Auth</a>
		</div>
		<div class="flex-none">
			<ul class="menu menu-horizontal px-1">
				<li>
					<form method="post" action="?/logout" use:enhance>
						<button type="submit">Logout</button>
					</form>
				</li>
			</ul>
		</div>
	</nav>
</header>

<div class="h-screen flex justify-center items-center">
	<h1 class="text-center text-3xl font-bold my-10">
		{data.username}
	</h1>
</div>
