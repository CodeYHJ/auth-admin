<script setup>
import { onMounted, ref } from 'vue'

import { getTokenList } from "@api"

import { notification } from 'ant-design-vue';

// import FormModal from "./modal.vue"

const columns = [
    {
        title: 'ID',
        dataIndex: 'id',
        key: 'id',
        align: "center",
    },
    {
        title: 'Code',
        dataIndex: 'code',
        key: 'code',
        align: "center"
    },
    {
        title: 'Access',
        key: 'access',
        dataIndex: 'access',
        align: "center"

    },
    {
        title: 'Refresh',
        key: 'refresh',
        dataIndex: 'refresh',
        align: "center"

    },
    {
        title: 'CreatedAt',
        key: 'created_at',
        dataIndex: 'created_at',
        align: "center"

    },
    {
        title: 'UpdatedAt',
        key: 'updated_at',
        dataIndex: 'updated_at',
        align: "center"
    },
    {
        title: '操作',
        key: 'op',
        width: 200,
        align: "center"

    },
]
const userid = ref(0)
const dataSource = ref([])
const visible = ref(false)
const showModal = (record) => {
    userid.value = record.id
    visible.value = true
}
const closeModal = () => {
    visible.value = false
}
const getData = () => {
    getTokenList().then(res => {
        dataSource.value = res
    })
}
const handleSubmit = (e) => {
    editAccount({ id: e.id, paw: e.newPass }).then(res => {
        notification.success({
            message: '修改密码成功'
        });
        getData()
    })
};
onMounted(() => {
    getData()
})
</script>

<template>
    <a-table :columns="columns" :data-source="dataSource" size="small">
        <template #bodyCell="{ column, record }">
            <template v-if="column.key === 'op'">
                <span>
                    <a @click="showModal(record)">Edit</a>
                    <a-divider type="vertical" />
                    <a>Delete</a>
                </span>
            </template>
        </template>
    </a-table>
    <!-- <FormModal :visible="visible" :id="userid" @onCancel="closeModal" @onSubmit="handleSubmit" /> -->

</template>

<style>
</style>
