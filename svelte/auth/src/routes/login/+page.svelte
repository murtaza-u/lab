<script lang="ts">
	import { enhance } from '$app/forms'
	import type { ActionData } from './$types'

	export let form: ActionData
	let loading = false
</script>

<h1 class="text-5xl font-bold text-center my-10">Login</h1>

<div class="mx-3 flex flex-col justify-center items-center">
	<form
		class="w-full lg:w-2/3"
		method="POST"
		use:enhance={() => {
			loading = true
			return async ({ update }) => {
				await update()
				loading = false
			}
		}}
	>
		<label class="form-control w-full my-3">
			<div class="label">
				<span class="label-text">Username</span>
				<span class="label-text-alt">*</span>
			</div>
			<input
				required
				name="username"
				type="text"
				placeholder="Enter your username"
				class="input input-bordered w-full"
				class:input-error={form?.ctx === 'username' || form?.ctx === 'global'}
			/>
			{#if form?.ctx === 'username'}
				<div class="label">
					<span class="label-text-alt text-error">{form.error}</span>
				</div>
			{/if}
		</label>
		<label class="form-control w-full my-3">
			<div class="label">
				<span class="label-text">Password</span>
				<span class="label-text-alt">*</span>
			</div>
			<input
				required
				name="password"
				type="password"
				placeholder="********"
				class="input input-bordered w-full"
				class:input-error={form?.ctx === 'password' || form?.ctx === 'global'}
			/>
			{#if form?.ctx === 'password'}
				<div class="label">
					<span class="label-text-alt text-error">{form.error}</span>
				</div>
			{/if}
		</label>
		<button type="submit" class="btn btn-secondary float-right my-3">
			Login
			<span
				class="loading loading-spinner loading-md"
				class:hidden={!loading}
			/>
		</button>
		{#if form?.ctx === 'global'}
			<div class="text-error my-3">{form.error}</div>
		{/if}
		<div></div>
	</form>

	<div class="mt-3">
		Do not have an account?
		<a href="/signup" class="link link-hover text-[blue]">Sign Up</a>
	</div>
</div>
