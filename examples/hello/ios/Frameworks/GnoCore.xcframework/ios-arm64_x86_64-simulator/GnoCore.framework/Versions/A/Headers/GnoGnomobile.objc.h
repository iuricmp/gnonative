// Objective-C API for talking to github.com/gnolang/gnomobile/framework Go package.
//   gobind -lang=objc -prefix="Gno" github.com/gnolang/gnomobile/framework
//
// File is generated by gobind. Do not edit.

#ifndef __GnoGnomobile_H__
#define __GnoGnomobile_H__

@import Foundation;
#include "ref.h"
#include "Universe.objc.h"


@class GnoGnomobileBridge;
@class GnoGnomobileBridgeConfig;
@protocol GnoGnomobilePromiseBlock;
@class GnoGnomobilePromiseBlock;

@protocol GnoGnomobilePromiseBlock <NSObject>
- (void)callReject:(NSError* _Nullable)error;
- (void)callResolve:(NSString* _Nullable)reply;
@end

@interface GnoGnomobileBridge : NSObject <goSeqRefInterface> {
}
@property(strong, readonly) _Nonnull id _ref;

- (nonnull instancetype)initWithRef:(_Nonnull id)ref;
- (nullable instancetype)init:(GnoGnomobileBridgeConfig* _Nullable)config;
- (BOOL)close:(NSError* _Nullable* _Nullable)error;
- (NSString* _Nonnull)getSocketPath;
- (long)getTcpPort;
@end

@interface GnoGnomobileBridgeConfig : NSObject <goSeqRefInterface> {
}
@property(strong, readonly) _Nonnull id _ref;

- (nonnull instancetype)initWithRef:(_Nonnull id)ref;
- (nullable instancetype)init;
@property (nonatomic) NSString* _Nonnull rootDir;
@property (nonatomic) NSString* _Nonnull tmpDir;
@property (nonatomic) BOOL useTcpListener;
@property (nonatomic) BOOL useUdsListener;
@end

FOUNDATION_EXPORT GnoGnomobileBridge* _Nullable GnoGnomobileNewBridge(GnoGnomobileBridgeConfig* _Nullable config, NSError* _Nullable* _Nullable error);

FOUNDATION_EXPORT GnoGnomobileBridgeConfig* _Nullable GnoGnomobileNewBridgeConfig(void);

@class GnoGnomobilePromiseBlock;

@interface GnoGnomobilePromiseBlock : NSObject <goSeqRefInterface, GnoGnomobilePromiseBlock> {
}
@property(strong, readonly) _Nonnull id _ref;

- (nonnull instancetype)initWithRef:(_Nonnull id)ref;
- (void)callReject:(NSError* _Nullable)error;
- (void)callResolve:(NSString* _Nullable)reply;
@end

#endif
