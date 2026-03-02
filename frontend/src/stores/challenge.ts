import { defineStore } from 'pinia';
import { ref } from 'vue';
import type { Challenge, ContainerInfo } from '@/types/challenge';
import { getChallenges, getChallenge } from '@/api/challenge';

export const useChallengeStore = defineStore('challenge', () => {
  // State
  const challenges = ref<Challenge[]>([]);
  const currentChallenge = ref<Challenge | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);

  // Actions
  async function fetchChallenges(params?: {
    category?: string;
    difficulty?: string;
    search?: string;
  }) {
    loading.value = true;
    error.value = null;

    try {
      const response = await getChallenges(params);
      challenges.value = response.items || [];
      return response;
    } catch (err: any) {
      error.value = err?.msg || 'Failed to fetch challenges';
      console.error('Fetch challenges error:', err);
      return null;
    } finally {
      loading.value = false;
    }
  }

  async function fetchChallenge(id: number) {
    loading.value = true;
    error.value = null;

    try {
      const response = await getChallenge(id);
      currentChallenge.value = response;
      return response;
    } catch (err: any) {
      error.value = err?.msg || 'Failed to fetch challenge';
      console.error('Fetch challenge error:', err);
      return null;
    } finally {
      loading.value = false;
    }
  }

  function setCurrentChallenge(challenge: Challenge | null) {
    currentChallenge.value = challenge;
  }

  function updateContainerState(challengeId: number, state: 'idle' | 'loading' | 'running', info?: ContainerInfo) {
    const challenge = challenges.value.find((c: Challenge) => c.id === challengeId);
    if (challenge) {
      challenge.containerState = state;
      if (info) {
        challenge.containerInfo = info;
      }
    }

    if (currentChallenge.value?.id === challengeId) {
      currentChallenge.value.containerState = state;
      if (info) {
        currentChallenge.value.containerInfo = info;
      }
    }
  }

  return {
    // State
    challenges,
    currentChallenge,
    loading,
    error,
    // Actions
    fetchChallenges,
    fetchChallenge,
    setCurrentChallenge,
    updateContainerState,
  };
});
