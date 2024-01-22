<script>
	import { page } from '$app/stores';
	import { Navbar, NavBrand, NavLi, NavUl, NavHamburger } from 'flowbite-svelte';
	$: activeUrl = $page.url.pathname;
	let activeClass = 'text-violet-950 md:text-white text-base  drop-shadow-lg font-bold';
	let nonActiveClass =
		'md:text-white hover:text-blue-800 text-base text-blue-700 font-bold md:border-0';
</script>

<Navbar class="bg-blue-base w-full">
	<NavBrand href="/">
		<img src="/assets/images/logo.png" class="me-3 h-12 sm:h-12" alt="Logo" />
		<span class="self-center whitespace-nowrap text-xl font-semibold text-white">QR Go</span>
	</NavBrand>
	<NavHamburger />
	<NavUl {activeUrl} {activeClass} {nonActiveClass}>
		<NavLi href="/">Home</NavLi>
		<NavLi href="/docs/components/navbar">How to</NavLi>
		<NavLi href="/docs/components/accordion">About Us</NavLi>
	</NavUl>
</Navbar>
