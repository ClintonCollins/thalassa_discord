// @generated by protoc-gen-connect-es v0.10.1 with parameter "target=ts"
// @generated from file thalassa/v1/thalassa.proto (package thalassa.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetCurrentSongPlayingRequest, GetCurrentSongPlayingResponse, GetSongRequestsRequest, GetSongRequestsResponse } from "./thalassa_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service thalassa.v1.APIService
 */
export const APIService = {
  typeName: "thalassa.v1.APIService",
  methods: {
    /**
     * @generated from rpc thalassa.v1.APIService.GetSongRequests
     */
    getSongRequests: {
      name: "GetSongRequests",
      I: GetSongRequestsRequest,
      O: GetSongRequestsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * rpc AddSongRequest(AddSongRequestRequest) returns (AddSongRequestResponse);
     *
     *  rpc SongRequestsUpdateStream(SongRequestsUpdateStreamRequest) returns (stream SongRequestsUpdateStreamResponse);
     *
     * @generated from rpc thalassa.v1.APIService.GetCurrentSongPlaying
     */
    getCurrentSongPlaying: {
      name: "GetCurrentSongPlaying",
      I: GetCurrentSongPlayingRequest,
      O: GetCurrentSongPlayingResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
