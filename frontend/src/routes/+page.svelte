<script lang="ts">
	import { AppShell } from '@svelteuidev/core';
	import { PUBLIC_BASE_API_URL } from '$env/static/public';

	import { Button } from 'flowbite-svelte';
	import ColorPicker from '../components/ColorPicker.svelte';
	import QrTypeButton from '../components/QrTypeButton.svelte';
	import { qrcodeContent } from '../stores/QrcodeStore';

	let qrImageUrl = '';

	const generateQRCode = async () => {
		try {
			if (!$qrcodeContent.data) {
				return 0;
			}
			const response = await fetch(`${PUBLIC_BASE_API_URL}`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					data: $qrcodeContent.data,
					fgColor: $qrcodeContent.fgColor,
					bgColor: $qrcodeContent.bgColor
				})
			});
			const qrCodeImage = await response.blob();
			qrImageUrl = URL.createObjectURL(qrCodeImage);
		} catch (error) {
			console.error('Error generating QR code:', error);
		}
	};

	const downloadFile = (url: string, fileName: string): void => {
		const anchor = document.createElement('a');
		anchor.href = url;
		anchor.download = fileName;

		document.body.appendChild(anchor);
		anchor.click();

		document.body.removeChild(anchor);
	};

	const handleChangeFgColor = (selectedColor: string): void => {
		$qrcodeContent.fgColor = selectedColor;
	};
	const handleChangeBgColor = (selectedColor: string): void => {
		$qrcodeContent.bgColor = selectedColor;
	};
</script>

<AppShell class="bg-blue-base">
	<div class="mt-16 flex justify-center pl-2 pr-2 md:mt-32 md:pl-32 md:pr-32">
		<div class="w-full max-w-6xl flex-row rounded-lg bg-white p-4 md:flex">
			<div class="flex w-full flex-col md:w-2/3">
				<QrTypeButton />
				{#if $qrcodeContent.fgColor === $qrcodeContent.bgColor}
					<p class="w-full text-center font-bold text-red-700">
						Caution! Avoid using the same color for QR codes.
					</p>
				{/if}
				<div class="mt-auto flex w-full flex-row self-end">
					<div class="w-1/2">
						<ColorPicker
							title={'Background Color'}
							changeColor={handleChangeBgColor}
							initColor={$qrcodeContent.bgColor}
						/>
					</div>
					<div class="w-1/2">
						<ColorPicker
							title={'Foreground Color'}
							changeColor={handleChangeFgColor}
							initColor={$qrcodeContent.fgColor}
						/>
					</div>
				</div>
			</div>

			<div class="bg-blue-light mt-6 justify-center md:mt-0 md:flex md:w-1/3 md:flex-col">
				<img
					src={`${qrImageUrl ? qrImageUrl : 'assets/images/qr-code-mockup.png'}`}
					alt="qr-code-mockup.png"
					class="w-full p-4"
				/>
				<div class="flex w-full p-3">
					<Button
						size="xs"
						class="text-blue-base border-blue-base mr-2  w-1/2 border-2 bg-white focus-within:ring-transparent hover:border-sky-700 hover:bg-sky-700 hover:text-white"
						on:click={() => {
							generateQRCode();
						}}>Create QR Code</Button
					><Button
						size="xs"
						class="mr-2 w-1/2  bg-teal-950  focus-within:ring-transparent hover:bg-teal-600"
						on:click={() => {
							downloadFile(qrImageUrl, 'qrcode.png');
						}}>Download</Button
					>
				</div>
			</div>
		</div>
	</div>
</AppShell>
