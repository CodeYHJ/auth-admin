<script setup>
import { toRefs, ref, reactive, defineProps, defineEmits } from 'vue'
const props = defineProps({
    visible: Boolean,
    id: Number
})
const emit = defineEmits(["onCancel", "onSubmit"])
const { visible, id } = toRefs(props)
const formRef = ref();
const formState = reactive({
    secret: '',
    domain: '',
});
const layout = {
    labelCol: { span: 4 },
    wrapperCol: { span: 14 },
};

const rules = {
    secret: [{
        required: true,
        trigger: 'change',
    }],
    domain: [{
        required: true,
        trigger: 'change',
    }]
};
const handleFinish = values => {
    values.id = id.value
    emit("onSubmit", values)
    emit("onCancel")
    resetForm()

};
const resetForm = () => {
    formRef.value.resetFields();
};
const handleModalCancel = () => {
    resetForm()
    emit("onCancel")
}
</script>

<template>
    <a-modal v-model:visible="visible" title="编辑" :footer="null" @cancel="handleModalCancel">
        <a-form ref="formRef" name="custom-validation" :model="formState" :rules="rules" v-bind="layout"
            @finish="handleFinish">
            <a-form-item has-feedback label="Secret" name="secret">
                <a-input v-model:value="formState.secret" type="text" autocomplete="off" />
            </a-form-item>
            <a-form-item has-feedback label="Domain" name="domain">
                <a-input v-model:value="formState.domain" type="text" autocomplete="off" />
            </a-form-item>
            <a-form-item :wrapper-col="{ span: 14, offset: 4 }">
                <a-button type="primary" html-type="submit">Submit</a-button>
                <a-button style="margin-left: 10px" @click="resetForm">Reset</a-button>
            </a-form-item>
        </a-form>
    </a-modal>
</template>