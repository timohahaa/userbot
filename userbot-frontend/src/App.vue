<script setup lang="ts">
import Logo from './components/Logo.vue'
import Accounts from './components/accounts/Accounts.vue'
import Channels from './components/channels/Channels.vue'
</script>

<template>
	<div class="z-10 container mx-auto h-[15vh] w-full">
		<Logo :size="100" />
	</div>
	<div class="z-10 container mx-auto grid grid-cols-2 gap-x-4 h-[70vh] w-full">
		<Channels />
		<Accounts />
	</div>
	<div class="lines">
		<span style="--i: 11;"></span>
		<span style="--i: 12;"></span>
		<span style="--i: 24;"></span>
		<span style="--i: 10;"></span>
		<span style="--i: 14;"></span>
		<span style="--i: 20;"></span>
		<span style="--i: 19;"></span>
		<span style="--i: 15;"></span>
		<span style="--i: 20;"></span>
		<span style="--i: 17;"></span>
		<span style="--i: 19;"></span>
		<span style="--i: 12;"></span>
		<span style="--i: 22;"></span>
		<span style="--i: 23;"></span>
		<span style="--i: 15;"></span>
		<span style="--i: 20;"></span>
		<span style="--i: 17;"></span>
		<span style="--i: 12;"></span>
		<span style="--i: 10;"></span>
		<span style="--i: 20;"></span>
		<span style="--i: 15;"></span>
		<span style="--i: 16;"></span>
		<span style="--i: 20;"></span>
		<span style="--i: 11;"></span>
		<span style="--i: 12;"></span>
		<span style="--i: 16;"></span>
	</div>
</template>
