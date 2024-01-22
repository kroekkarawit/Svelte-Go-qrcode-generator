<script lang="ts">
	import { qrcodeContent } from '../../stores/QrcodeStore';

	let inputText: string = '';
	$: qrCodeContent = $qrcodeContent;
</script>

<div class="mt-4 flex w-full flex-col p-1">
	<p class="text-sm font-bold">Your Text</p>
	<textarea
		class="border-sky-light w-full rounded-lg border-2 bg-[#eaf3ff]"
		bind:value={inputText}
		on:input={() => ($qrcodeContent.data = inputText)}
	></textarea>
</div>
