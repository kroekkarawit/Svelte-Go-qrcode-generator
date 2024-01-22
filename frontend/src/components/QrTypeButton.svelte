<script lang="ts">
	import Text from '../components/icons/Text.svelte';
	import Crypto from '../components/icons/Crypto.svelte';
	import Email from '../components/icons/Email.svelte';
	import Event from './icons/Event.svelte';
	import Phone from '../components/icons/Phone.svelte';
	import SMS from '../components/icons/SMS.svelte';
	import Url from '../components/icons/Url.svelte';
	import Wifi from '../components/icons/Wifi.svelte';
	import { Button } from 'flowbite-svelte';

	import QrText from './inputGroups/QrText.svelte';
	import QrUrl from './inputGroups/QrUrl.svelte';
	import QrEmail from './inputGroups/QrEmail.svelte';
	import QrPhone from './inputGroups/QrPhone.svelte';
	import QrSMS from './inputGroups/QrSMS.svelte';
	import QrWifi from './inputGroups/QrWifi.svelte';
	import QrEvent from './inputGroups/QrEvent.svelte';
	import QrCrypto from './inputGroups/QrCrypto.svelte';

	let qrcodeType: { name: string; icon: any; form: any }[] = [
		{
			name: 'Text',
			icon: Text,
			form: QrText
		},
		{
			name: 'URL',
			icon: Url,
			form: QrUrl
		},
		{
			name: 'Email',
			icon: Email,
			form: QrEmail
		},
		{
			name: 'Phone',
			icon: Phone,
			form: QrPhone
		},
		{
			name: 'SMS',
			icon: SMS,
			form: QrSMS
		},
		{
			name: 'Wifi',
			icon: Wifi,
			form: QrWifi
		},
		{
			name: 'Event',
			icon: Event,
			form: QrEvent
		},
		{
			name: 'Crypto',
			icon: Crypto,
			form: QrCrypto
		}
	];

	let selectQRCodeType: string = 'Text';

	$: selectedForm = qrcodeType.find((i) => i.name === selectQRCodeType)?.form;
</script>

<div class="w-full">
	<div class="lg:grid-row-1 grid grid-flow-col grid-rows-2 gap-1 md:grid-rows-2 md:gap-2">
		{#each qrcodeType as { name, icon }}
			<Button
				size="xs"
				class={`mr-2 bg-teal-950 outline-none focus-within:ring-white hover:bg-teal-600 ${selectQRCodeType === name && ' bg-sky-200 text-teal-950'}`}
				on:click={() => (selectQRCodeType = name)}
			>
				<svelte:component
					this={icon}
					classes={` me-2 h-3 w-3 ${selectQRCodeType === name ? 'fill-teal-900' : 'fill-white'}`}
				/>
				{name}
			</Button>
		{/each}
	</div>
	{#if selectedForm}
		<svelte:component this={selectedForm} />
	{/if}
</div>
