#ifndef MOCHIGO_H
#define MOCHIGO_H

#include <stdbool.h>

GoRef matchaGoBool(bool);
bool matchaGoToBool(GoRef);
GoRef matchaGoInt(int);
GoRef matchaGoInt64(int64_t);
int64_t matchaGoToInt64(GoRef);
GoRef matchaGoUint64(uint64_t);
uint64_t matchaGoToUint64(GoRef);
GoRef matchaGoFloat64(double);
double matchaGoToFloat64(GoRef);
GoRef matchaGoString(CGoBuffer); // Frees the buffer
CGoBuffer matchaGoToString(GoRef);
GoRef matchaGoBytes(CGoBuffer); // Frees the buffer
CGoBuffer matchaGoToBytes(GoRef);

GoRef matchaGoArray();
int64_t matchaGoArrayLen(GoRef);
GoRef matchaGoArrayAppend(GoRef, GoRef);
GoRef matchaGoArrayAt(GoRef, int64_t);

GoRef matchaGoMap();
GoRef matchaGoMapKeys(GoRef);
GoRef matchaGoMapGet(GoRef map, GoRef key);
void matchaGoMapSet(GoRef map, GoRef key, GoRef value);

GoRef matchaGoType(CGoBuffer); // Frees the buffer
GoRef matchaGoFunc(CGoBuffer); // Frees the buffer

bool matchaGoIsNil(GoRef);
bool matchaGoEqual(GoRef, GoRef);
GoRef matchaGoElem(GoRef);
GoRef matchaGoCall(GoRef, CGoBuffer, GoRef);
GoRef matchaGoField(GoRef, CGoBuffer);
void matchaGoFieldSet(GoRef, CGoBuffer, GoRef);

void matchaGoUntrack(GoRef);

#import <Foundation/Foundation.h>
@class MatchaGoValue;

@interface MatchaGoBridge : NSObject
+ (MatchaGoBridge *)sharedBridge;
@end

@interface MatchaGoValue : NSObject
//- (id)initWithGoRef:(GoRef)ref;
- (id)initWithObject:(id)v;
- (id)initWithBool:(BOOL)v;
- (id)initWithInt:(int)v;
- (id)initWithLongLong:(long long)v;
- (id)initWithUnsignedLongLong:(unsigned long long)v;
- (id)initWithDouble:(double)v;
- (id)initWithString:(NSString *)v;
- (id)initWithData:(NSData *)v;
- (id)initWithArray:(NSArray<MatchaGoValue *> *)v;
- (id)initWithType:(NSString *)typeName;
- (id)initWithFunc:(NSString *)funcName;
@property (nonatomic, readonly) GoRef ref;
- (id)toObject;
- (BOOL)toBool;
- (long long)toLongLong;
- (unsigned long long)toUnsignedLongLong;
- (double)toDouble;
- (NSString *)toString;
- (NSData *)toData;
- (NSArray *)toArray;
- (BOOL)isNil;
- (BOOL)isEqual:(MatchaGoValue *)value;
- (MatchaGoValue *)elem;
- (NSArray<MatchaGoValue *> *)call:(NSString *)method, ... NS_REQUIRES_NIL_TERMINATION; // pass in nil for the method to call a closure.
- (NSArray<MatchaGoValue *> *)call:(NSString *)method args:(va_list)args;
- (MatchaGoValue *)field:(NSString *)name;
- (void)setField:(NSString *)name value:(MatchaGoValue *)value;
- (MatchaGoValue *)objectForKeyedSubscript:(NSString *)key;
- (void)setObject:(MatchaGoValue *)object forKeyedSubscript:(NSString *)key;
@end

#endif // MOCHIGO_H
