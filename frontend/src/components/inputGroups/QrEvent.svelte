<script lang="ts">
	import { qrcodeContent } from '../../stores/QrcodeStore';

	let inputTitle: string = '';
	let inputStartTime: string = '';
	let inputEndTime: string = '';
	$: qrCodeContent = $qrcodeContent;

	const formatDateToCustomFormat = (date: Date): string => {
		const year = date.getFullYear();
		const month = (date.getMonth() + 1).toString().padStart(2, '0');
		const day = date.getDate().toString().padStart(2, '0');
		const hours = date.getHours().toString().padStart(2, '0');
		const minutes = date.getMinutes().toString().padStart(2, '0');
		const seconds = date.getSeconds().toString().padStart(2, '0');

		return `${year}${month}${day}T${hours}${minutes}${seconds}`;
	};
</script>

<div class="mt-4 grid w-full grid-cols-1 gap-4 p-1">
	<div class="w-full">
		<p class="text-sm font-bold">Event Title</p>
		<input
			type="text"
			class="border-sky-light w-full rounded-lg bg-[#eaf3ff]"
			bind:value={inputTitle}
			on:change={() =>
				($qrcodeContent.data = `BEGIN:VEVENT\nSUMMARY:${inputTitle}\nLOCATION:\nDTSTART:${formatDateToCustomFormat(new Date(inputStartTime))}\nDTEND:${formatDateToCustomFormat(new Date(inputEndTime))}\nEND:VEVENT`)}
		/>
	</div>
	<div class="w-full">
		<p class="text-sm font-bold">Start time</p>
		<input
			type="datetime-local"
			class="border-sky-light w-full rounded-lg bg-[#eaf3ff]"
			bind:value={inputStartTime}
			on:change={() =>
				($qrcodeContent.data = `BEGIN:VEVENT\nSUMMARY:${inputTitle}\nLOCATION:\nDTSTART:${formatDateToCustomFormat(new Date(inputStartTime))}\nDTEND:${formatDateToCustomFormat(new Date(inputEndTime))}\nEND:VEVENT`)}
		/>
	</div>
	<div class="w-full">
		<p class="text-sm font-bold">End time</p>
		<input
			type="datetime-local"
			class="border-sky-light w-full rounded-lg bg-[#eaf3ff]"
			bind:value={inputEndTime}
			on:change={() =>
				($qrcodeContent.data = `BEGIN:VEVENT\nSUMMARY:${inputTitle}\nLOCATION:\nDTSTART:${formatDateToCustomFormat(new Date(inputStartTime))}\nDTEND:${formatDateToCustomFormat(new Date(inputEndTime))}\nEND:VEVENT`)}
		/>
	</div>
</div>
