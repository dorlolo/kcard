export interface ApiError { code: string; message: string; details?: string[] }
export interface ApiClientOptions { baseUrl?: string; getAuthToken?: () => string | undefined; getWorkspaceId?: () => string | undefined }

export class ApiClient {
  constructor(private readonly options: ApiClientOptions = {}) {}

  async request<T>(path: string, init: RequestInit = {}): Promise<T> {
    const headers = new Headers(init.headers)
    headers.set('Content-Type', 'application/json')
    headers.set('X-Request-ID', crypto.randomUUID())
    const token = this.options.getAuthToken?.()
    const workspaceID = this.options.getWorkspaceId?.() ?? '00000000-0000-0000-0000-000000000001'
    if (token) headers.set('Authorization', `Bearer ${token}`)
    headers.set('X-Workspace-ID', workspaceID)
    const response = await fetch(`${this.options.baseUrl ?? import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080/api/v1'}${path}`, { ...init, headers })
    if (!response.ok) throw await normalizeError(response)
    return response.json() as Promise<T>
  }
}

async function normalizeError(response: Response): Promise<ApiError> {
  try { const body = await response.json(); return body.error ?? { code: 'request_failed', message: response.statusText } } catch { return { code: 'request_failed', message: response.statusText } }
}

export const apiClient = new ApiClient()
