<template>
	<div class="bg-table-bg-color flex-col w-full h-full rounded-lg px-6 py-4">
		<div class="flex justify-between items-center mb-[24px] h-[64px]">
			<div class="flex flex-col h-full">
				<h2 class="text-lg text-text-color mb-1">Accounts</h2>
				<span class="text-gray-400 text-[12px]"
					>More than 400+ new accounts</span
				>
			</div>
			<button
				class="bg-button-color h-[46px] flex flex-row items-center font-medium rounded px-4"
			>
				<div class="h-[24px] w-[24px] mr-2">
					<img :src="account" class="max-h-[100%]" />
				</div>
				<span class="text-button-text-color text-xs">New Account</span>
			</button>
		</div>
		<div class="w-full h-[calc(100%-64px-64px-24px-24px)] mb-[24px]">
			<table :class="tableClasses">
				<tr
					class="text-left h-[42px] text-gray-400 text-[14px] border-checkbox-color border-b-2"
				>
					<th class="pb-4">
						<div class="h-full flex items-center justify-start">
							<label class="custom-checkbox">
								<input type="checkbox" />
								<span class="checkmark"></span>
							</label>
						</div>
					</th>
					<th class="font-medium pr-3 pb-4">NAME</th>
					<th class="font-medium pr-3 pb-4">ID</th>
					<th class="text-right font-medium pb-4">ACTION</th>
				</tr>
				<tr
					v-if="accounts.length != null"
					v-for="account in accounts.slice(currentPage * 4, (currentPage + 1) * 4)"
					:class="trClasses"
				>
					<td class="w-[12px] pr-3">
						<!-- в этом div pt-1 это пиксель дрочь, чтобы ровно было -->
						<div class="h-full flex items-center justify-start">
							<label class="custom-checkbox">
								<input type="checkbox" />
								<span class="checkmark"></span>
							</label>
						</div>
					</td>
					<td class="pr-3 py-2 text-text-color">
						{{ account.name }}
					</td>
					<td class="pr-3 py-2 text-text-color">{{ account.id }}</td>
					<td class="flex justify-end items-center h-full">
						<div class="mr-2">
							<button
								class="h-[36px] w-[36px] bg-checkbox-color p-1 flex justify-center items-center rounded"
							>
								<img :src="edit" class="h-[20px] w-[20px]" />
							</button>
						</div>
						<button
							class="h-[36px] w-[36px] bg-checkbox-color p-1 flex justify-center items-center rounded"
						>
							<img :src="trash" class="h-[20px] w-[20px]" />
						</button>
					</td>
				</tr>
				<tr></tr>
			</table>
		</div>
		<div class="flex flex-row justify-center items-center h-[64px]">
			<button
				class="h-[16px] w-[16px]"
				@click="
					() => {
						if (currentPage !== 0) {
							currentPage--
						}
					}
				"
			>
				<img :src="leftArrow" />
			</button>
			<div
				class="h-[32px] w-[32px] flex justify-center items-center text-text-color border-white border rounded text-xs font-normal mx-2"
			>
				{{ currentPage }}
			</div>
			<button
				class="h-[16px] w-[16px]"
				@click="
					() => {
						if (accounts.length > (currentPage + 1) * 4) {
							currentPage++
						}
					}
				"
			>
				<img :src="rightArrow" />
			</button>
		</div>
	</div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue'
import { Account } from './account'
import axios from 'axios'
import config from '../../config.ts'
import account from '../../assets/add-account.png' // не думаю, что так норм импортить изображения
import edit from '../../assets/edit.png'
import trash from '../../assets/trash.png'
import leftArrow from '../../assets/left-arrow.png'
import rightArrow from '../../assets/right-arrow.png'

let accounts = ref<Array<Account>>([])
let currentPage = ref(0)
console.log(window)

async function fetchAccounts(): Promise<void> {
	// фиг знает написал в этой функции от себя, вообще хз, что здесь нужно
	let resp = await axios.get(`${config.BASE_URL}/accountnel`)
	for (let ac of resp.data.data) {
		let account = new Account(ac.accountnel_id, ac.accountnel_name)
		accounts.value?.push(account)
	}
}

onMounted(() => {
	const acs: Account[] = [
		new Account(1, 'GG WP'),
		new Account(2, 'reaper'),
		new Account(3, 'reaper'),
		new Account(4, 'reaper'),
		new Account(5, 'reaper'),
		new Account(6, 'reaper'),
	]
	// await fetchaccountnels()
	accounts.value?.push(...acs)
	console.log(accounts.value?.length)
})

const tableClasses = computed(() => {
	// расчёт выделяемой
	const length = accounts.value?.slice(
		currentPage.value * 4,
		(currentPage.value + 1) * 4
	).length
	console.log(length)
	let tableHeight = ''

	if (length === 0) {
		tableHeight = 'h-[42px]'
	} else if (length === 1) {
		tableHeight = 'h-[calc((100%-42px)*0.25+42px)]'
	} else if (length === 2) {
		tableHeight = 'h-[calc((100%-42px)*0.5+42px)]'
	} else if (length === 3) {
		tableHeight = 'h-[calc((100%-42px)*0.75+42px)]'
	} else {
		tableHeight = 'h-full'
	}
	return `w-full border-collapse m-0 p-0 ${tableHeight}`
})

const trClasses = computed(() => {
	const length = accounts.value?.slice(
		currentPage.value * 4,
		(currentPage.value + 1) * 4
	).length
	console.log(length)
	let trHeight = ''

	if (length === 1) {
		trHeight = 'h-full'
	} else if (length === 2) {
		trHeight = 'h-[calc(100%*0.5)]'
	} else if (length === 3) {
		trHeight = 'h-[calc(100%*0.33)]'
	} else {
		trHeight = 'h-[calc(100%*0.25)]'
	}
	return `${trHeight} border-checkbox-color border-b-2`
})
</script>
