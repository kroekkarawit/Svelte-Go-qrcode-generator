<script lang="ts">
	import { qrcodeContent } from '../../stores/QrcodeStore';

	let inputPhoneNumber: string = '';
	let inputMessage: string = '';

	$: qrCodeContent = $qrcodeContent;
</script>

<div class="mt-4 grid w-full grid-cols-1 gap-4 p-1">
	<div class="w-full">
		<p class="text-sm font-bold">Your Phone Number</p>
		<input
			type="text"
			class="border-sky-light w-full rounded-lg bg-[#eaf3ff]"
			bind:value={inputPhoneNumber}
			on:change={() => ($qrcodeContent.data = `SMSTO:${inputPhoneNumber}:${inputMessage}`)}
		/>
	</div>
	<div class="w-full">
		<p class="text-sm font-bold">Your Message</p>
		<textarea
			class="border-sky-light w-full rounded-lg bg-[#eaf3ff]"
			bind:value={inputMessage}
			on:change={() => ($qrcodeContent.data = `SMSTO:${inputPhoneNumber}:${inputMessage}`)}
		></textarea>
	</div>
</div>
