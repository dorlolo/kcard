export function confirmDestructiveAction(message: string): boolean { return window.confirm(message) }
export function undoLabel(action: string): string { return `Undo ${action}` }
