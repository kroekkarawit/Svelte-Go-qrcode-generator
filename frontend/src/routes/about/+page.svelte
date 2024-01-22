<script>
	let hiddenButton;
	let colorInput;

	function autoClick() {
		hiddenButton.click();
	}

	function handleColorChange(event) {
		console.log('Selected Color:', event.target.value);
	}
</script>

<button on:click={autoClick} bind:this={hiddenButton}>Click me</button>
<input type="color" bind:value={colorInput} on:change={handleColorChange} />

<style>
	/* Initially hide the color input */
	input {
		display: none;
	}
</style>
