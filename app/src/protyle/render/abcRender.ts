import {addScript} from "../util/addScript";
import {Constants} from "../../constants";

declare const ABCJS: {
    renderAbc(element: Element, text: string, options: { responsive: string }): void;
};

export const abcRender = (element: Element, cdn = Constants.PROTYLE_CDN) => {
    let abcElements: Element[] = [];
    if (element.getAttribute("data-subtype") === "abc") {
        // 编辑器内代码块编辑渲染
        abcElements = [element];
    } else {
        abcElements = Array.from(element.querySelectorAll('[data-subtype="abc"]'));
    }
    if (abcElements.length === 0) {
        return;
    }
    if (abcElements.length > 0) {
        addScript(`${cdn}/js/abcjs/abcjs-basic-min.js?v=6.2.2`, "protyleAbcjsScript").then(() => {
            abcElements.forEach((e: HTMLDivElement) => {
                if (e.getAttribute("data-render") === "true") {
                    return;
                }
                if(!e.firstElementChild.classList.contains("protyle-icons")) {
                    e.insertAdjacentHTML("afterbegin", '<div class="protyle-icons"><span class="protyle-icon protyle-icon--first protyle-action__edit"><svg><use xlink:href="#iconEdit"></use></svg></span><span class="protyle-icon protyle-action__menu protyle-icon--last"><svg><use xlink:href="#iconMore"></use></svg></span></div>');
                }
                if (e.childElementCount < 4) {
                    e.lastElementChild.insertAdjacentHTML("beforebegin", `<span style="position: absolute">${Constants.ZWSP}</span>`);
                }
                const renderElement = e.firstElementChild.nextElementSibling as HTMLElement;
                ABCJS.renderAbc(renderElement, Lute.UnEscapeHTMLStr(e.getAttribute("data-content")), {
                    responsive: "resize"
                });
                renderElement.setAttribute("contenteditable", "false");
                e.setAttribute("data-render", "true");
            });
        });
    }
};
