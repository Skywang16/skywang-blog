import { createApp, nextTick } from 'vue';
import wLoading from './wLoading.vue';

const Mask = createApp(wLoading).mount(document.createElement('div'));

// 主函数
const TzLoading = {
  install: (app) => {
    // 蒙版操作更新
    const toggleLoading = function (el, binding) {
      if (binding.value) {
        nextTick(() => {
          el.instance.loading = binding.value.loading || false;
          if (binding.modifiers.fullscreen) {
            // 全屏情况下
            document.body.appendChild(el.mask);
          } else {
            el.appendChild(el.mask);
          }
        });
      } else {
        // 移除节点
        el.mask && el.mask.parentNode && el.mask.parentNode.removeChild(el.mask);
        el.instance && el.instance.$unmount();
      }
    };

    app.directive('wLoading', {
      // 只调用一次，指令第一次绑定到元素时调用。在这里可以进行一次性的初始化设置。
      mounted(el, binding) {
        let loading = binding.value.loading;

        const instance = Mask;
        instance.loading = loading || false;

        el.instance = instance; // 将instance存入
        el.mask = instance.$el; // dom存入，方便获取
        toggleLoading(el, binding);
      },
      // 所在组件的 VNode 更新时调用
      updated(el, binding) {
        if (binding.oldValue !== binding.value) {
          toggleLoading(el, binding);
        }
      },
      // 只调用一次，指令与元素解绑时调用
      beforeUnmount(el) {
        try {
          el.mask && el.mask.parentNode && el.mask.parentNode.removeChild(el.mask);
          el.instance && el.instance.$unmount();
        } catch (e) {
          console.log(e);
        }
      },
    });
  },
};

export default TzLoading;