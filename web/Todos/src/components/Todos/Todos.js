import {reactive, ref} from '@vue/reactivity'
import Button from "../Button/Button.vue"
import Task from "../Task/Task.vue"
import {useToast} from "vue-toastification";
import axios from 'axios';


export default {
    components: {
        Button,
        Task
    },

    setup() {
        let missions = reactive([])
        const input = ref(null)
        const filter = ref("all")
        const task_edit = ref({
            index: 0,
            text: " ",
            select_task_for_edit: false
        })

        const toast = useToast()

        axios.defaults.baseURL = 'http://localhost:8075';

        function Add() {
            if (input.value.value !== "") {
                axios.post('/api/task', {name: input.value.value})
                    .then(response => {
                        missions.push(response.data);
                        input.value.value = "";
                    })
                    .catch(error => {
                        toast.error(error.message);
                    });
            } else {
                toast.error("Поле задачи не может быть пустым");
            }
        }

        function Edit() {
            if (task_edit.value.select_task_for_edit) {
                axios.put(`/api/task/${task_edit.value.index}`, {name: input.value.value})
                    .then(response => {
                        let mission = missions.find(mission => mission.id === task_edit.value.index);
                        if (mission) {
                            mission.name = input.value.value;
                        }
                        task_edit.value.select_task_for_edit = false;
                        input.value.value = "";
                        toast.success(response.data);
                    })
                    .catch(error => {
                        toast.error(error.message);
                    });
            }
        }


        function ChangeCondition(index) {
            axios.put(`/api/task/${index}/complete`)
                .then(response => {
                    toast.success("Статус задачи изменен");
                })
                .catch(error => {
                    toast.error(error.message);
                });
        }

        function Delete(index) {
            axios.delete(`/api/task/${index}`)
                .then(response => {
                    toast.success("Задача удалена");
                })
                .catch(error => {
                    toast.error(error.message);
                });
        }

        axios.get('/api/tasks')
            .then(response => {
                if (response.data !== null) {
                    response.data.forEach(task => {
                        missions.push(task)
                    });
                }
            })
            .catch(error => {
                toast.error(error.message);
            });


        return {
            missions,
            Add,
            Edit,
            ChangeCondition,
            Delete,
            input,
            toast,
            task_edit,
            filter
        }
    }
}
