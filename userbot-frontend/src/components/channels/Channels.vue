<template>
    <div class="grid grid-cols-1 grid-rows-2 text-base text-text max-w-fit">
        <div class="flex items-center justify-center font-bold">
            Registered channels
        </div>
        <div>
            <div class="grid grid-cols-3">
                <div class="mx-2">Name</div>
                <div class="mx-2">Id</div>
                <div class="mx-2">Subscribers</div>
            </div>
            <div v-for="chan of chans" class="grid grid-cols-3">
                <div class="mx-2">{{ chan.name }}</div>
                <div class="mx-2">{{ chan.id }}</div>
                <div class="mx-2">{{ chan.userount }}</div>
            </div>
        </div>
        <div>
            <button @click="fetchChannels" class="btn bg-dark-main text-base">Refresh</button>
        </div>
    </div>
    <Add/>
</template>


<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { Channel } from './channel';
import  Add  from './Add.vue'
import axios from 'axios';
import config from "../../config.ts"


let chans = ref<Array<Channel>>([])

async function fetchChannels(): Promise<void> {
    let resp = await axios.get(`${config.BASE_URL}/channel`)
    for(let ch of resp.data.data) {
        let chan = new Channel(ch.channel_id, ch.user_count, ch.channel_name)
        chans.value?.push(chan)
    }
}

onMounted( () => {
    const chs: Channel[] = [
        new Channel(1, 5, "reaper 121313212"),
        new Channel(2, 424234234, "reaper 2"),
        new Channel(3, 7, "reaper 3"),
    ]
    // await fetchChannels()
    chans.value?.push(...chs)

    console.log(chans.value?.length)
})

</script>