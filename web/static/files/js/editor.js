document.addEventListener('DOMContentLoaded', () => {
    // === Получаем конфигурацию из глобальной переменной ===
    const config = window.gonote_config || {};
    const {
        noteID = 'new',
        isRussian = false,
        initialContent = '',
        initialTitle = '',
        initialAuthor = ''
    } = config;

    // === DOM-элементы ===
    const charCounter = document.getElementById('charCounter');
    const titleInput = document.querySelector('.title-input');
    const authorInput = document.querySelector('.meta-input');
    const pubBtn = document.getElementById('pubBtn');
    const passBtn = document.getElementById('passBtn');
    const pubModal = document.getElementById('pubModal');
    const confirmPubBtn = document.getElementById('confirmPub');
    const msgBox = document.getElementById('msgBox');

    if (!charCounter || !titleInput || !authorInput || !pubBtn || !passBtn) {
        console.error('Editor: required elements not found');
        return;
    }

    // === Константы ===
    const LIMIT_NOTE = 1000000;
    const LIMIT_MENU = 10000;
    const STORAGE_KEY = `gonote-editor-content-${noteID}`;

    // === Иконки Quill ===
    const icons = Quill.import('ui/icons');
    icons['image-link'] = `
    <svg viewBox="0 0 18 18" xmlns="http://www.w3.org/2000/svg">
      <line class="ql-stroke" x1="3" y1="4" x2="15" y2="4"></line>
      <line class="ql-stroke" x1="3" y1="4" x2="3" y2="14"></line>
      <line class="ql-stroke" x1="15" y1="4" x2="15" y2="7"></line>
      <line class="ql-stroke" x1="3" y1="14" x2="7" y2="14"></line>
      <circle class="ql-fill" cx="6" cy="7" r="1"></circle>
      <g transform="scale(0.8) translate(7,7)">
        <line class="ql-stroke" x1="7" x2="11" y1="7" y2="11"></line>
        <path class="ql-even ql-stroke" d="M8.9,4.577a3.476,3.476,0,0,1,.36,4.679A3.476,3.476,0,0,1,4.577,8.9C3.185,7.5,2.035,6.4,4.217,4.217S7.5,3.185,8.9,4.577Z"></path>
        <path class="ql-even ql-stroke" d="M13.423,9.1a3.476,3.476,0,0,0-4.679-.36,3.476,3.476,0,0,0,.36,4.679c1.392,1.392,2.5,2.542,4.679.36S14.815,10.5,13.423,9.1Z"></path>
      </g>
    </svg>`;

    icons['hr'] = `
    <svg viewBox="0 0 18 18" xmlns="http://www.w3.org/2000/svg">
      <line class="ql-stroke" x1="3" x2="15" y1="9" y2="9"></line>
    </svg>`;

    // === Регистрация HR Blot (только один раз) ===
    const BlockEmbed = Quill.import('blots/block/embed');
    if (!Quill.imports['blots/hr']) {
        class HrBlot extends BlockEmbed {}
        HrBlot.blotName = 'hr';
        HrBlot.tagName = 'hr';
        Quill.register(HrBlot);
    }

    // === Инициализация Quill ===
    const quill = new Quill('#editor', {
        theme: 'bubble',
        modules: {
            toolbar: {
                container: [
                    ['bold', 'italic', 'underline', 'strike'],
                    ['blockquote', 'code-block'],
                    ['link', 'image', 'image-link', 'video', 'formula'],
                    [{ 'list': 'ordered' }, { 'list': 'bullet' }, { 'list': 'check' }],
                    [{ 'script': 'sub' }, { 'script': 'super' }],
                    [{ 'indent': '-1' }, { 'indent': '+1' }],
                    [{ 'direction': 'rtl' }],
                    [{ 'header': [1, 2, 3, 4, 5, 6, false] }],
                    [{ 'size': ['small', false, 'large', 'huge'] }],
                    [{ 'color': [] }, { 'background': [] }],
                    [{ 'font': [] }],
                    [{ 'align': [] }],
                    ['hr'],
                    ['clean']
                ],
                handlers: {
                    'hr': function () {
                        const range = this.quill.getSelection(true);
                        if (range && range.length > 0) {
                            quill.deleteText(range.index, range.length);
                        }
                        quill.insertEmbed(range.index, 'hr', true, Quill.sources.USER);
                        quill.setSelection(range.index + 1);
                    },
                    'image-link': function () {
                        const tooltip = this.quill.theme.tooltip;
                        tooltip.edit('image', '');
                        tooltip.textbox.placeholder = 'Image Url';

                        const originalSave = tooltip.save;
                        tooltip.save = function () {
                            const url = tooltip.textbox.value;
                            if (url) {
                                const range = quill.getSelection(true);
                                if (range && range.length > 0) {
                                    quill.deleteText(range.index, range.length);
                                }
                                quill.insertEmbed(range.index, 'image', url);
                                quill.setSelection(range.index + 1);
                            }
                            tooltip.hide();
                        };
                    },
                    'clean': function () {
                        const range = this.quill.getSelection();
                        if (!range || range.length === 0) return;

                        const delta = this.quill.getContents(range.index, range.length);
                        let replaced = false;

                        delta.ops.forEach((op, i) => {
                            if (op.insert && op.insert.image) {
                                const imageUrl = op.insert.image;
                                this.quill.deleteText(range.index + i, 1, 'user');
                                this.quill.insertText(range.index + i, imageUrl, { link: imageUrl }, 'user');
                                replaced = true;
                            }
                        });

                        if (!replaced) {
                            this.quill.removeFormat(range.index, range.length, 'user');
                        }
                    }
                }
            }
        },
        placeholder: 'Your note...'
    });

    const menuQuill = new Quill('#menu-editor', {
        theme: 'bubble',
        modules: {
            toolbar: {
                container: [
                    ['bold', 'italic'],
                    [{'header': [2, 3, false]}],
                    [{ 'align': [] }],
                    ['link'],
                    ['hr'],
                    ['clean']
                ],
                handlers: {
                    'hr': function () {
                        const range = this.quill.getSelection(true);
                        if (range && range.length > 0) {
                            menuQuill.deleteText(range.index, range.length);
                        }
                        menuQuill.insertEmbed(range.index, 'hr', true, Quill.sources.USER);
                        menuQuill.setSelection(range.index + 1);
                    }
                }
            }
        },
        placeholder: 'Menu (links, headers...)'
    });

    menuQuill.on('selection-change', (range) => {
        const menuContainer = document.getElementById('menu-editor');
        if (menuContainer) {
            if (range) {
                // Фокус есть
                menuContainer.style.borderColor = '#007aff';
                menuContainer.style.background = '#fff';
            } else {
                // Фокус потерян
                menuContainer.style.borderColor = '#ccc';
                menuContainer.style.background = '#fafafa';
                menuContainer.style.boxShadow = 'none';
            }
        }
    });

    // === Восстановление из localStorage ===
    const saved = localStorage.getItem(STORAGE_KEY);
    if (saved) {
        try {
            const data = JSON.parse(saved);
            titleInput.value = data.title || initialTitle;
            authorInput.value = data.author || initialAuthor;
            if (data.content) quill.root.innerHTML = data.content;
            if (data.menu) menuQuill.root.innerHTML = data.menu;
        } catch (e) {
            console.warn('Failed to load saved draft', e);
        }
    } else {
        titleInput.value = initialTitle;
        authorInput.value = initialAuthor;
        if (initialContent) quill.root.innerHTML = initialContent;
    }

    // === Paste handler ===
    quill.root.addEventListener('paste', async (e) => {
        const clipboardData = e.clipboardData || window.clipboardData;
        const text = clipboardData.getData('text').trim();

        if (!text || !/^https?:\/\//i.test(text)) return;

        e.preventDefault();
        const range = quill.getSelection(true) || { index: quill.getLength() };
        const currentHost = window.location.origin;
        const siteRegex = new RegExp(
            `^${currentHost.replace(/[-\/\\^$*+?.()|[\]{}]/g, '\\$&')}/note/([a-zA-Z0-9_-]+)`,
            'i'
        );
        const match = text.match(siteRegex);

        if (match) {
            const noteID = match[1];
            try {
                const resp = await fetch(`${currentHost}/api/note/${noteID}`);
                if (!resp.ok) throw new Error('Not found');
                const data = await resp.json();
                const title = data.title || noteID;
                quill.insertText(range.index, title, { link: text });
                quill.setSelection(range.index + title.length);
            } catch {
                quill.insertText(range.index, text, { link: text });
                quill.setSelection(range.index + text.length);
            }
            return;
        }

        const img = new Image();
        img.onload = () => {
            quill.insertEmbed(range.index, 'image', text, 'user');
            quill.setSelection(range.index + 1);
        };
        img.onerror = () => {
            quill.insertText(range.index, text, { link: text });
            quill.setSelection(range.index + text.length);
        };
        img.src = text;
    });

    menuQuill.root.addEventListener('paste', async (e) => {
        const clipboardData = e.clipboardData || window.clipboardData;
        const text = clipboardData.getData('text').trim();

        // Пропускаем, если не URL
        if (!text || !/^https?:\/\//i.test(text)) return;

        e.preventDefault();

        const range = menuQuill.getSelection(true) || { index: menuQuill.getLength() };

        try {
            const encodedURL = encodeURIComponent(text);
            const resp = await fetch(`/api/getlinktitle?url=${encodedURL}`);
            const data = await resp.json();
            const title = data.title || text; // если прокси вернул пусто — fallback на URL

            menuQuill.insertText(range.index, title, { link: text }, Quill.sources.USER);
            menuQuill.setSelection(range.index + title.length);
        } catch (err) {
            console.warn('Failed to fetch title, using URL as fallback', err);
            // Вставляем как есть
            menuQuill.insertText(range.index, text, { link: text }, Quill.sources.USER);
            menuQuill.setSelection(range.index + text.length);
        }
    });

    // === Счётчик символов ===
    function updateCounter() {
        const len = quill.root.innerHTML.length;
        charCounter.textContent = `Limit note size: ${len} / ${LIMIT_NOTE}`;
        charCounter.style.color = len > LIMIT_NOTE ? 'red' : 'gray';
        charCounter.style.opacity = len > LIMIT_NOTE ? 0 : Math.min(len / LIMIT_NOTE, 1.0);
    }

    // === Автопрокрутка ===
    function autoScrollOnEdit(delta, oldDelta, source) {
        if (source !== 'user') return;
        const range = quill.getSelection();
        if (!range) return;
        const bounds = quill.getBounds(range.index);
        const bottom = bounds.top + window.scrollY + bounds.height;
        const shouldScroll = bottom > window.scrollY + window.innerHeight - 100;
        if (shouldScroll) {
            window.scrollBy({ top: bounds.height + 20, behavior: 'smooth' });
        }
    }

    // === Сохранение в localStorage ===
    function saveAll() {
        const data = {
            title: titleInput.value,
            author: authorInput.value,
            content: quill.root.innerHTML,
            menu: menuQuill.root.innerHTML
        };
        localStorage.setItem(STORAGE_KEY, JSON.stringify(data));
    }

    // === Публикация ===
    function pubNote(password = '') {
        const title = titleInput.value.trim();
        const author = authorInput.value.trim();
        const content = quill.root.innerHTML;
        const menu = menuQuill.root.innerHTML;

        if (!title) return GonoteUtils.showMessage(isRussian ? 'Заголовок не может быть пустым' : 'Title cannot be empty');
        if (title.length < 3) return GonoteUtils.showMessage(isRussian ? 'Заголовок слишком короткий' : 'Title is too small');
        if (!content || quill.root.innerText.trim().length === 0) return GonoteUtils.showMessage(isRussian ? 'Контент не может быть пустым' : 'Content cannot be empty');
        if (content.length > LIMIT_NOTE) return GonoteUtils.showMessage(isRussian ? 'Превышен лимит размера!' : 'Content exceeds maximum size!');
        if (menu.length > LIMIT_MENU) return GonoteUtils.showMessage(isRussian ? 'Превышен лимит размера меню!' : 'Menu content exceeds maximum size!');

        const link = noteID === 'new' ? '/new' : `/edit/${noteID}`;

        fetch(link, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ title, author, content, menu, password })
        })
            .then(res => res.json())
            .then(data => {
                if (data.ok) {
                    localStorage.removeItem(STORAGE_KEY);
                    window.location.href = `/note/${data.noteID}`;
                } else {
                    GonoteUtils.showMessage(data.error || 'Unknown error');
                }
            })
            .catch(err => {
                console.error(err);
                GonoteUtils.showMessage('Network error');
            });
    }

    // === Обработчики событий ===
    quill.on('text-change', (delta, oldDelta, source) => {
        updateCounter();
        saveAll();
        autoScrollOnEdit(delta, oldDelta, source);
    });

    menuQuill.on('text-change', (delta, oldDelta, source) => {
        saveAll();
    });

    titleInput.addEventListener('input', saveAll);
    authorInput.addEventListener('input', saveAll);

    pubBtn.addEventListener('click', () => pubNote(''));

    const preGeneratedPassword = GonoteUtils.generatePassword();

    passBtn.addEventListener('click', () => {
        document.getElementById('modalPassword').innerText = preGeneratedPassword;
        pubModal.style.display = 'flex';
    });

    confirmPubBtn.addEventListener('click', () => {
        pubNote(preGeneratedPassword);
        pubModal.style.display = 'none';
    });

    pubModal.addEventListener('click', (e) => {
        if (e.target === pubModal) pubModal.style.display = 'none';
    });

    // === Инициализация ===
    updateCounter();
});