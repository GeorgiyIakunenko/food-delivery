<script setup>
import {
  Dialog,
  DialogPanel,
  DialogTitle,
  TransitionChild,
  TransitionRoot,
} from "@headlessui/vue";
import {
  ExclamationTriangleIcon,
  ExclamationCircleIcon,
  CheckCircleIcon,
} from "@heroicons/vue/24/outline";

const { type, open, title } = defineProps({
  open: {
    type: Boolean,
    default: false,
  },
  type: {
    type: String,
    validator(value) {
      return ["warning", "success", "error"].includes(value);
    },
    default: "warning",
  },
  title: {
    type: String,
    default: "Error",
  },
});

const emits = defineEmits(["modalClose"]);

const closeModal = (value) => {
  emits("modalClose", value);
};
</script>

<template>
  <TransitionRoot as="template" :show="open">
    <Dialog as="div" class="relative z-10" @close="closeModal(false)">
      <TransitionChild
        as="template"
        enter="ease-out duration-300"
        enter-from="opacity-0"
        enter-to="opacity-100"
        leave="ease-in duration-200"
        leave-from="opacity-100"
        leave-to="opacity-0"
      >
        <div
          class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"
        />
      </TransitionChild>

      <div class="fixed inset-0 z-10 overflow-y-auto">
        <div
          class="flex min-h-full items-center justify-center p-4 text-center sm:items-center sm:p-0"
        >
          <TransitionChild
            as="template"
            enter="ease-out duration-300"
            enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
            enter-to="opacity-100 translate-y-0 sm:scale-100"
            leave="ease-in duration-200"
            leave-from="opacity-100 translate-y-0 sm:scale-100"
            leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
            class="w-4/5"
          >
            <DialogPanel
              class="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg"
            >
              <div class="bg-white px-4 pb-4 pt-5 sm:p-6 sm:pb-4">
                <div class="sm:flex sm:items-start">
                  <div
                    class="mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full sm:mx-0 sm:h-10 sm:w-10"
                  >
                    <ExclamationTriangleIcon
                      v-if="type === 'warning'"
                      class="text-red-600"
                      aria-hidden="true"
                    />
                    <CheckCircleIcon
                      v-if="type === 'success'"
                      class="text-green-600"
                      aria-hidden="true"
                    />
                    <ExclamationCircleIcon
                      v-if="type === 'error'"
                      class="text-red-600"
                      aria-hidden="true"
                    />
                  </div>
                  <div class="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
                    <DialogTitle
                      as="h3"
                      class="text-base font-semibold capitalize leading-6 text-gray-900"
                      >{{ title }}</DialogTitle
                    >
                    <div class="mt-2">
                      <p class="text-sm text-gray-500">
                        <slot></slot>
                      </p>
                    </div>
                  </div>
                </div>
              </div>
              <div
                class="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6"
              >
                <button
                  v-if="type === 'warning'"
                  type="button"
                  class="inline-flex w-full justify-center rounded-md bg-red-600 px-3 py-2 text-sm font-semibold text-white shadow-sm transition-all duration-300 hover:bg-red-500 sm:ml-3 sm:w-auto"
                  @click="closeModal(true)"
                >
                  Yes
                </button>
                <button
                  v-if="type === 'warning'"
                  type="button"
                  class="mt-3 inline-flex w-full justify-center rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 transition-all duration-300 hover:bg-gray-50 sm:mt-0 sm:w-auto"
                  @click="closeModal(false)"
                  ref="cancelButtonRef"
                >
                  No
                </button>
                <button
                  v-if="type === 'error'"
                  type="button"
                  class="mt-3 inline-flex w-full justify-center rounded-md bg-neutral-40 px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 transition-all duration-300 hover:bg-gray-50 sm:mt-0 sm:w-auto"
                  @click="closeModal(true)"
                  ref="cancelButtonRef"
                >
                  Ok
                </button>
                <button
                  v-if="type === 'success'"
                  type="button"
                  class="mt-3 inline-flex w-full justify-center rounded-md bg-green-400 px-3 py-2 text-sm font-semibold text-neutral-0 shadow-sm ring-1 ring-inset transition-all duration-300 hover:bg-green-200 hover:text-neutral-900 sm:mt-0 sm:w-auto"
                  @click="closeModal(true)"
                  ref="cancelButtonRef"
                >
                  Ok
                </button>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>
