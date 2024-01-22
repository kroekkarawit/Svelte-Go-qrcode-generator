<script lang="ts">
	import { qrcodeContent } from '../../stores/QrcodeStore';

	let inputEmail: string = '';
	let inputSubject: string = '';
	let inputMessage: string = '';

	$: qrCodeContent = $qrcodeContent;
</script>

<div class="mt-4 grid w-full grid-cols-1 gap-4 p-1">
	<div class="w-full">
		<p class="text-sm font-bold">Your Email</p>
		<input
			type="email"
			class="border-sky-light w-full rounded-lg bg-[#eaf3ff]"
			bind:value={inputEmail}
			placeholder="name@email.com"
			on:change={() =>
				($qrcodeContent.data = `mailto:${inputEmail}?subject=${inputSubject}&body=${inputMessage}`)}
		/>
	</div>
	<div class="w-full">
		<p class="text-sm font-bold">Subject</p>
		<input
			type="text"
			class="border-sky-light w-full rounded-lg bg-[#eaf3ff]"
			bind:value={inputSubject}
			on:change={() =>
				($qrcodeContent.data = `mailto:${inputEmail}?subject=${inputSubject}&body=${inputMessage}`)}
		/>
	</div>

	<div class="w-full">
		<p class="text-sm font-bold">Message</p>
		<textarea
			class="border-sky-light w-full rounded-lg bg-[#eaf3ff]"
			bind:value={inputMessage}
			on:input={() =>
				($qrcodeContent.data = `mailto:${inputEmail}?subject=${inputSubject}&body=${inputMessage}`)}
		></textarea>
	</div>
</div>
