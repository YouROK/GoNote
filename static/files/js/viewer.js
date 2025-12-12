document.addEventListener('DOMContentLoaded', () => {
    // === Конфигурация из шаблона ===
    const config = window.gonote_config || {};
    const {
        noteID = '',
        hasEdit = false,
        hasPass = false,
        isRussian = false
    } = config;

    // === Инициализация Quill (readonly) ===
    const BlockEmbed = Quill.import('blots/block/embed');
    if (!Quill.imports['blots/hr']) {
        class HrBlot extends BlockEmbed {}
        HrBlot.blotName = 'hr';
        HrBlot.tagName = 'hr';
        Quill.register(HrBlot);
    }

    const editorEl = document.getElementById('editor');
    if (editorEl) {
        new Quill('#editor', {
            readOnly: true,
            theme: 'bubble'
        });
    }
    const menuEl = document.getElementById('menu-editor');
    if (editorEl) {
        new Quill('#menu-editor', {
            readOnly: true,
            theme: 'bubble'
        });
    }

    // === Кнопка "Copy Link" ===
    const copyBtn = document.getElementById('copyBtn');
    if (copyBtn) {
        copyBtn.addEventListener('click', async () => {
            try {
                await navigator.clipboard.writeText(window.location.href);
                const originalText = copyBtn.textContent;
                copyBtn.textContent = isRussian ? 'Скопировано!' : 'Copied!';
                setTimeout(() => {
                    copyBtn.textContent = originalText;
                }, 1500);
            } catch (err) {
                console.error('Failed to copy:', err);
                copyBtn.textContent = isRussian ? 'Ошибка' : 'Error';
                setTimeout(() => {
                    copyBtn.textContent = isRussian ? 'Копировать ссылку' : 'Copy Link';
                }, 1500);
            }
        });
    }

    // === Кнопка "Share" (Web Share API) ===
    const shareBtn = document.getElementById('shareBtn');
    if (shareBtn && navigator.share) {
        shareBtn.addEventListener('click', async () => {
            const title = document.querySelector('.header-title')?.textContent || 'GoNote';
            const text = isRussian ? 'Заметка на GoNote' : 'Note on GoNote';
            try {
                await navigator.share({ title, text, url: window.location.href });
            } catch (err) {
                if (err.name !== 'AbortError') {
                    console.warn('Share failed:', err);
                }
            }
        });
    } else if (shareBtn) {
        // Скрываем кнопку, если Web Share не поддерживается
        shareBtn.style.display = 'none';
    }

    // === Модальное окно: Report ===
    const reportBtn = document.getElementById('reportBtn');
    const reportModal = document.getElementById('reportModal');
    const reportReason = document.getElementById('reportReason');
    const reportEmail = document.getElementById('reportEmail');
    const reportText = document.getElementById('reportText');
    const submitReport = document.getElementById('submitReport');

    if (reportBtn && reportModal) {
        reportBtn.addEventListener('click', () => {
            reportModal.style.display = 'flex';
            if (reportReason) reportReason.focus();
        });

        const isValidEmail = (email) => {
            const parts = email.split('@');
            return parts.length === 2 && parts[0] && parts[1];
        };

        if (submitReport) {
            submitReport.addEventListener('click', async () => {
                const reason = reportReason?.value?.trim();
                const email = reportEmail?.value?.trim();
                const text = reportText?.value?.trim();

                if (!reason || !text) {
                    GonoteUtils.showMessage(isRussian ? 'Заполните все поля' : 'Please fill in all fields');
                    return;
                }
                if (!email || !isValidEmail(email)) {
                    GonoteUtils.showMessage(isRussian ? 'Неверный email' : 'Please enter a valid email');
                    return;
                }

                try {
                    const res = await fetch('/report', {
                        method: 'POST',
                        headers: { 'Content-Type': 'application/json' },
                        body: JSON.stringify({ reason, text, link: window.location.href, email })
                    });

                    const data = await res.json();
                    if (data.status === 'ok') {
                        GonoteUtils.showMessage(isRussian ? 'Жалоба отправлена!' : 'Complaint sent!');
                        reportModal.style.display = 'none';
                        if (reportReason) reportReason.value = '';
                        if (reportText) reportText.value = '';
                    } else {
                        GonoteUtils.showMessage(isRussian ? 'Ошибка отправки' : 'Error sending complaint');
                    }
                } catch (err) {
                    console.error(err);
                    GonoteUtils.showMessage(isRussian ? 'Ошибка сети' : 'Network error');
                }
            });
        }

        // Закрытие по клику вне модалки или Esc
        const closeModal = () => {
            reportModal.style.display = 'none';
        };

        reportModal.addEventListener('click', (e) => {
            if (e.target === reportModal) closeModal();
        });

        document.addEventListener('keydown', (e) => {
            if (reportModal.style.display === 'flex' && e.key === 'Escape') {
                closeModal();
            }
        });
    }

    // === Модальное окно: Edit with Password ===
    const editBtn = document.getElementById('editBtn');
    const passwordModal = document.getElementById('passwordModal');
    const passwordInput = document.getElementById('notePassword');
    const submitPassword = document.getElementById('submitPassword');

    if (editBtn && (hasEdit || hasPass)) {
        editBtn.addEventListener('click', () => {
            if (hasEdit) {
                window.location.href = `/note/${noteID}/edit`;
            } else if (hasPass) {
                passwordModal.style.display = 'flex';
                if (passwordInput) passwordInput.focus();
            }
        });

        if (submitPassword && passwordInput) {
            submitPassword.addEventListener('click', () => {
                const password = passwordInput.value;
                if (!password) return;

                fetch(`/note/${noteID}/checkpass`, {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ password })
                })
                    .then(res => res.json())
                    .then(data => {
                        if (data.ok) {
                            window.location.href = `/note/${noteID}/edit`;
                        } else {
                            GonoteUtils.showMessage(isRussian ? 'Неверный пароль' : 'Incorrect password');
                        }
                    })
                    .catch(err => {
                        console.error(err);
                        GonoteUtils.showMessage(isRussian ? 'Ошибка проверки' : 'Verification error');
                    });
            });

            passwordInput.addEventListener('keypress', (e) => {
                if (e.key === 'Enter') submitPassword.click();
            });
        }

        // Закрытие модалки пароля
        if (passwordModal) {
            passwordModal.addEventListener('click', (e) => {
                if (e.target === passwordModal) {
                    passwordModal.style.display = 'none';
                    if (passwordInput) passwordInput.value = '';
                }
            });

            document.addEventListener('keydown', (e) => {
                if (passwordModal.style.display === 'flex' && e.key === 'Escape') {
                    passwordModal.style.display = 'none';
                    if (passwordInput) passwordInput.value = '';
                }
            });
        }
    }
});