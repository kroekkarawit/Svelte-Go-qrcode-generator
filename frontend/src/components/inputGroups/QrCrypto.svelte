<script lang="ts">
	import { qrcodeContent } from '../../stores/QrcodeStore';
	import { Radio } from 'flowbite-svelte';

	let inputCoin: string = '';
	let inputAddress: string = '';
	let inputAmount: string = '';

	$: qrCodeContent = $qrcodeContent;

	const handleRadioChange = (event: Event | null): void => {
		if (event) {
			const target = event.target as HTMLInputElement;
			inputCoin = target.value;
		}
	};
</script>

<div class="mt-4 grid w-full grid-cols-1 gap-4 p-1">
	<div class="w-full">
		<p class="text-sm font-bold">Coin</p>
		<ul
			class="w-full items-center divide-x divide-gray-200 rounded-lg border border-gray-200 sm:flex rtl:divide-x-reverse dark:divide-gray-600 dark:border-gray-600 dark:bg-gray-800"
		>
			<li class="w-full">
				<Radio
					name="hor-list"
					class="p-3"
					bind:group={inputCoin}
					value="bitcoin"
					on:change={handleRadioChange}>Bitcoin</Radio
				>
			</li>
			<li class="w-full">
				<Radio
					name="hor-list"
					class="p-3"
					bind:group={inputCoin}
					value="bitcoincash"
					on:change={handleRadioChange}>Bitcoin Cash</Radio
				>
			</li>
			<li class="w-full">
				<Radio
					name="hor-list"
					class="p-3"
					bind:group={inputCoin}
					value="ethereum"
					on:change={handleRadioChange}>Ethereum</Radio
				>
			</li>
			<li class="w-full">
				<Radio
					name="hor-list"
					class="p-3"
					bind:group={inputCoin}
					value="litecoin"
					on:change={handleRadioChange}>Litecoin</Radio
				>
			</li>
		</ul>
	</div>
	<div class="w-full">
		<p class="text-sm font-bold">Address</p>
		<input
			type="text"
			class="border-sky-light w-full rounded-lg bg-[#eaf3ff]"
			bind:value={inputAddress}
			on:change={() => ($qrcodeContent.data = `bitcoin:${inputAddress}?amount=${inputAmount}`)}
		/>
	</div>
	<div class="w-full">
		<p class="text-sm font-bold">Amount</p>
		<input
			type="number"
			class="border-sky-light w-full rounded-lg bg-[#eaf3ff]"
			bind:value={inputAmount}
			on:change={() => ($qrcodeContent.data = `bitcoin:${inputAddress}?amount=${inputAmount}`)}
		/>
	</div>
</div>
