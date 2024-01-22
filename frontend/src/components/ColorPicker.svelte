<script lang="ts">
	import { colors } from '@svelteuidev/core';
	import { qrcodeContent } from '../stores/QrcodeStore';
	import { onMount } from 'svelte';

	export let title: string;
	let selectedColor: string;
	let listsColor: string[] = [
		'#000000',
		'#ffffff',
		'#81ffd5',
		'#8a2be1',
		'#ff4500',
		'#008001',
		'#0302fc',
		'#ffbfcb',
		'#dfff01',
		'#fff9e3',
		'#e2d912',
		'#d9138a',
		'#10a4d8',
		'#322e2f',
		'#e62166',
		'#0e1037',
		'#1e3d59',
		'#f6f0e3',
		'#ff6e41',
		'#ffc13c',
		'#edc19c',
		'#1d8480',
		'#e5e5db',
		'#d9a5b2',
		'#c8d9ed',
		'#781e47',
		'#7a1f4a',
		'#1f2862',
		'#8b307f',
		'#ff9a8c',
		'#e7f0f9',
		'#188afc',
		'#00ebff'
	];
	export let changeColor: (color: string) => void;
	export let initColor: string;

	onMount(() => {
		selectedColor = initColor || listsColor[0]; // Use initColor if provided, otherwise default to the first color
	});
</script>

<div class="flex flex-col p-2">
	<p class="text-center text-lg font-bold">{title}</p>
	<p class="text-md border-2 border-blue-100 text-center font-bold text-blue-600">
		{selectedColor}
	</p>
	<div class="mt-2 grid grid-flow-col grid-rows-7 gap-y-1 md:grid-rows-3">
		{#each listsColor as color}
			<button
				class={`h-4 w-4 p-3 md:h-2 md:w-2 md:p-2 ${selectedColor === color ? 'border-2 border-gray-400' : 'border-2 border-white'}`}
				style="background-color: {color}; border-radius: 50%;"
				on:click={() => {
					selectedColor = color;
					changeColor(selectedColor);
				}}
			></button>
		{/each}
	</div>
</div>
<div class="w-1/2"></div>
