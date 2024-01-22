<script lang="ts">
	import { qrcodeContent } from '../../stores/QrcodeStore';

	let inputText: string = '';
	$: qrCodeContent = $qrcodeContent;
</script>

<div class="mt-4 flex w-full flex-col p-1">
	<p class="text-sm font-bold">Your Phone Number</p>
	<input
		type="text"
		class="border-sky-light w-full rounded-lg bg-[#eaf3ff]"
		bind:value={inputText}
		on:change={() => ($qrcodeContent.data = inputText)}
	/>
</div>
