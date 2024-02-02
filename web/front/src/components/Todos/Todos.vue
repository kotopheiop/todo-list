<template>
  <div class="flex justify-center items-baseline h-full pt-4 sm:p-16">
    <div class="w-[30rem] h-full m-4 flex flex-col">
      <div class="w-full mb-12 flex justify-between items-center">
        <div>
          <i class="fa-solid fa-list-ul mr-4 sm:text-2xl align-baseline text-gray-400"></i>
          <h1 class="text-sm sm:text-2xl font-bold inline-block text-gray-400">To Do List</h1>
        </div>
        <div>
          <ul :class="{'all' : filter === 'all' , 'done' : filter === 'done' , 'undone' : filter === 'undone'}"
              class="flex justify-between relative overflow-hidden afte:block after:h-[40px] sm:after:h-[45px] after:bg-slate-900 after:absolute after:rounded-lg after:-z-10 after:transition-all">
            <li @click="filter = 'all'" class="py-2 px-4 sm:text-xl text-gray-400 cursor-pointer">All</li>
            <li @click="filter = 'done'" class="p-2 px-4 sm:text-xl text-gray-400 cursor-pointer">Done</li>
            <li @click="filter = 'undone'" class="p-2 px-4 sm:text-xl text-gray-400 cursor-pointer">Undone</li>
          </ul>
        </div>
      </div>

      <div class="w-full flex flex-row">
        <input @keypress.enter="task_edit.select_task_for_edit === false ? Add() : Edit()" ref="input"
               type="text" placeholder="Новая задача"
               class="input basis-3/4 bg-transparent border-0 border-b-2 rounded-none p-3 focus:outline-none border-b-slate-900 text-gray-400 placeholder:ttext-gray-400 shadow-2xl"/>

        <Button @click="Add" Value="Добавить" AddClass="basis-1/4 bg-slate-900 text-gray-400 hover:text-gray-400"
                v-show="task_edit.select_task_for_edit === false"/>
        <Button @click="Edit" Value="Редактировать" AddClass="basis-1/4 bg-slate-900 text-gray-400 hover:text-gray-400"
                v-show="task_edit.select_task_for_edit === true"/>
      </div>
      <div class="h-full mt-4">
        <Task v-for="(mission , index) in missions" :key="mission.id" :input="mission.name"
              @showFunc="missions.splice(index , 1), Delete(mission.id)"
              @changeCondition="mission.complete = !mission.complete, ChangeCondition(mission.id)"
              @edit="task_edit.index = mission.id , task_edit.select_task_for_edit = true , input.value = mission.name"
              :Condition="[{'line-through': mission.complete} , {'border-green-500': mission.complete}]"
              v-show="filter === 'all' ? true : (filter === 'done' ? mission.complete : !mission.complete)"
        />
      </div>
    </div>
  </div>
</template>

<script src="./Todos.js"></script>

<style src="./Todos.css"></style>