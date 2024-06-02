<template>
    <div class="w-full h-full flex justify-start items-end pb-2">
        <div class="h-[40px] w-[36px] pb-1 mr-2">
            <img :src="scsm_logo" class="max-h-[100%]"/>
        </div>
        <h1 class="text-text-color text-4xl">Sarcasm</h1>
    </div>
</template>


<script setup lang="ts">
import scsm_logo from "../assets/sarcasm_logo.png"

const props = defineProps<{
    size: number
}>()
</script>

<style scoped>
img {
    width: v-bind(size + 'px');
    height: v-bind(size + 'px');
}
</style>