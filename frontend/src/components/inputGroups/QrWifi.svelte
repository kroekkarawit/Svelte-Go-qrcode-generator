<script lang="ts">
	import { qrcodeContent } from '../../stores/QrcodeStore';

	let inputSSID: string = '';
	let inputPassword: string = '';
	let inputEncryption: string = 'nopass';

	$: qrCodeContent = $qrcodeContent;
</script>

<div class="mt-4 grid w-full grid-cols-1 gap-4 p-1">
	<div class="w-full">
		<p class="text-sm font-bold">Wireless SSID</p>
		<input
			type="text"
			class="border-sky-light w-full rounded-lg bg-[#eaf3ff]"
			bind:value={inputSSID}
			on:change={() =>
				($qrcodeContent.data = `WIFI:S:${inputSSID};T:${inputEncryption};P:${inputPassword};;`)}
		/>
	</div>
	<div class="w-full">
		<p class="text-sm font-bold">Password</p>
		<input
			type="text"
			class="border-sky-light w-full rounded-lg bg-[#eaf3ff]"
			bind:value={inputPassword}
			on:change={() =>
				($qrcodeContent.data = `WIFI:S:${inputSSID};T:${inputEncryption};P:${inputPassword};;`)}
		/>
	</div>
	<div class="w-full">
		<p class="text-sm font-bold">Encryption</p>
		<select
			class="border-sky-light w-full rounded-lg bg-[#eaf3ff]"
			bind:value={inputEncryption}
			on:change={() =>
				($qrcodeContent.data = `WIFI:S:${inputSSID};T:${inputEncryption};P:${inputPassword};;`)}
		>
			<option value="nopass" selected="selected">no encryption</option>
			<option value="WEP">WEP</option>
			<option value="WPA">WPA/WPA2</option>
		</select>
	</div>
</div>
